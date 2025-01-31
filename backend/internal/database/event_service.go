package database

import (
	"backend/internal/forms"
	"backend/internal/location"
	"backend/internal/models"
	"backend/pkg/normalizer"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type EventService interface {
	GetEvents(params map[string]any, limit int, preloadOptions []string) ([]*models.Event, error)
	CreateEvent(event forms.Event) (models.Event, error)
	DeleteEvent(id int) (models.Event, error)
	UpdateEvent() (models.Event, error)
	GetEvent(eventId uint) (*models.Event, error)
	PromoteEvent(id int) (*models.Event, error)
	ReportEvent(form forms.ReportForm) error
}

type eventServiceImpl struct {
	db *gorm.DB
}

func (e *eventServiceImpl) ReportEvent(form forms.ReportForm) error {
	var event models.Event
	if err := e.db.First(&event, "id = ?", form.ID).Error; err != nil {
		return fmt.Errorf("failed to find event: %w", err)
	}

	eventToReport := &models.ReportedEvent{
		Reason:  form.Reason,
		EventID: event.ID,
	}

	if err := e.db.Create(eventToReport).Error; err != nil {
		return fmt.Errorf("failed to create reported event: %w", err)
	}

	return nil
}

func (e *eventServiceImpl) PromoteEvent(id int) (*models.Event, error) {
	var event models.Event

	if err := e.db.First(&event, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("event not found")
	}

	event.IsPromoted = true

	if err := e.db.Save(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func NewEventService(db *gorm.DB) EventService {
	return &eventServiceImpl{
		db: db,
	}
}

func (s *service) EventService() EventService {
	return s.eventService
}

func (e *eventServiceImpl) CreateEvent(event forms.Event) (models.Event, error) {
	newEvent := models.Event{
		Title:       event.Title,
		Description: event.Description,
		ImageURL:    &event.ImageURL,
		IsAdultOnly: event.IsAdultOnly,
		EventType:   models.EventType(event.EventType),
		StartDate:   event.StartDate,
		FinishDate:  event.FinishDate,
	}

	event.Tags = normalizer.Normalizer(event.Tags)

	ch := location.FetchLocation(event.Lon, event.Lat)

	result := <-ch
	if result.Err != nil {
		log.Println(result.Err)
		return models.Event{}, result.Err
	}

	location := result.Location

	var organizers []*models.User

	e.db.Where("id IN ? AND email IS NOT NULL", event.Organizers).Find(&organizers)

	if len(organizers) == 0 {
		return models.Event{}, fmt.Errorf("no organizers")
	}

	tags, _ := getOrCreateTags(e.db, event.Tags)

	newEvent.Tags = tags
	newEvent.LocationID = location.ID
	newEvent.Location = &location
	newEvent.EventOrganizers = organizers
	newEvent.IsAdultOnly = event.IsAdultOnly

	e.db.Create(&newEvent)

	return newEvent, nil
}

func getOrCreateTags(tx *gorm.DB, tags []string) ([]*models.Tag, error) {
	var existingTags []*models.Tag
	tagNames := make(map[string]*models.Tag)

	if err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Tag{}).
			Where("name IN ?", tags).
			Find(&existingTags).Error; err != nil {
			return err
		}

		for _, tag := range existingTags {
			tagNames[tag.Name] = tag
		}

		return nil
	}); err != nil {
		return nil, err
	}

	var toCreate []*models.Tag
	for _, tag := range tags {
		if _, exists := tagNames[tag]; !exists {
			toCreate = append(toCreate, &models.Tag{
				Name: tag,
			})
		}
	}

	if len(toCreate) > 0 {
		if err := tx.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&toCreate).Error; err != nil {
				return err
			}

			for _, createdTag := range toCreate {
				tagNames[createdTag.Name] = createdTag
			}
			return nil
		}); err != nil {
			log.Printf("Failed to create tags: %v", err)
			return nil, err
		}
	}

	var result []*models.Tag
	for _, tag := range tagNames {
		result = append(result, tag)
	}

	return result, nil
}

func (e *eventServiceImpl) DeleteEvent(id int) (models.Event, error) { return models.Event{}, nil }
func (e *eventServiceImpl) UpdateEvent() (models.Event, error)       { return models.Event{}, nil }

type PreloadOption struct {
	Association string
}

/*
Filter Events by given params. Result length <= limit.

Params:

  - params: json key - value
  - limit: max. amount of events to find

Returns:

  - list of events
  - error occured during transaction
*/
func (e *eventServiceImpl) GetEvents(params map[string]any, limit int, preloadOptions []string) ([]*models.Event, error) {
	q := e.db.Model(&models.Event{})

	for _, option := range preloadOptions {
		q = q.Preload(option)
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	if params["lon"] != nil && params["lat"] != nil && params["accuracy"] != nil && params["lon"] != "" && params["lat"] != "" && params["accuracy"] != "" {
		log.Println(params["lon"], params["lat"], params["accuracy"])

		tmpQ := q.Joins("JOIN locations ON locations.id = events.location_id").
			Joins("JOIN coords ON coords.id = locations.coords_id").
			Where("ST_DWithin(coords.geom::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)",
				params["lon"], params["lat"], params["accuracy"])

		if q.Error != nil {
			log.Println(q.Error.Error())
		} else {
			delete(params, "lon")
			delete(params, "lat")
			delete(params, "accuracy")
			q = tmpQ
		}
	}

	for tag, value := range params {
		switch tag {
		case "title":
			if v, ok := value.(string); ok {
				q = q.Where("title LIKE ?", "%"+v+"%")
			}

		case "location", "city", "country":
			if v, ok := value.(string); ok {
				if tag == "location" {
					q = q.Joins("JOIN locations ON locations.id = events.location_id").
						Where("locations.name LIKE ?", "%"+v+"%")
				} else {
					q = q.Joins("JOIN locations ON locations.id = events.location_id").
						Where(fmt.Sprintf("locations.%s LIKE ?", tag), "%"+v+"%")
				}
			}

		case "street", "street_number":
			if v, ok := value.(string); ok {
				q = q.Joins("JOIN addresses ON addresses.location_id = locations.id").
					Where(fmt.Sprintf("addresses.%s LIKE ?", tag), "%"+v+"%")
			}

		case "event_tags":
			if v, ok := value.(string); ok {
				tags := strings.Split(v, ",")
				for _, tag := range tags {
					q = q.Joins("JOIN event_tags ON event_tags.event_id = events.id").
						Joins("JOIN tags ON tags.id = event_tags.tag_id").
						Where("tags.name = ?", tag)
				}
			}

		default:
			q = q.Where(fmt.Sprintf("%s = ?", tag), value)
		}
	}

	var res []*models.Event
	if err := q.Debug().Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (e *eventServiceImpl) GetEvent(eventId uint) (*models.Event, error) {
	var event models.Event
	if err := e.db.Preload("Location").
		Preload("Location.Address").
		Preload("Tags").
		Preload("EventOrganizers").
		Preload("Votes").
		First(&event, "id = ?", eventId).Error; err != nil {
		return nil, err
	}

	return &event, nil
}
