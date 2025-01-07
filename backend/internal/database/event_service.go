package database

import (
	"backend/internal/models"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type EventService interface {
	GetEvents(params map[string]any, limit int) ([]*models.Event, error)
	CreateEvent() (models.Event, error)
	DeleteEvent() (models.Event, error)
	UpdateEvent() (models.Event, error)
}

type eventServiceImpl struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) EventService {
	return &eventServiceImpl{
		db: db,
	}
}

func (s *service) EventService() EventService {
	return s.eventService
}

func (e *eventServiceImpl) CreateEvent() (models.Event, error) { return models.Event{}, nil }
func (e *eventServiceImpl) DeleteEvent() (models.Event, error) { return models.Event{}, nil }
func (e *eventServiceImpl) UpdateEvent() (models.Event, error) { return models.Event{}, nil }

/*
Filter Events by given params. Result length <= limit.

Params:

  - params: json key - value
  - limit: max. amount of events to find

Returns:

  - list of events
  - error occured during transaction
*/
func (e *eventServiceImpl) GetEvents(params map[string]any, limit int) ([]*models.Event, error) {
	q := e.db.
		Preload("Location").
		Preload("Location.Address").
		Preload("Tags").
		Model(&models.Event{})

	if limit > 0 {
		q = q.Limit(limit)
	}

	if params["lon"] != "" && params["lat"] != "" && params["accuracy"] != "" {
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
