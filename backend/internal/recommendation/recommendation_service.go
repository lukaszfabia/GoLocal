package recommendation

import (
	"backend/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type RecommendationService interface {
	Predict([]*models.Event, uint, int) ([]uint, error)
	ModifyAttendancePreference(uint, uint, bool) error
	GetUserPreferences(uint) (*models.UserPreference, error)
}

func findEventByID(events []*models.Event, eventID uint) (*models.Event, error) {
	for _, event := range events {
		if event.ID == eventID {
			return event, nil
		}
	}
	return nil, fmt.Errorf("event with ID %d not found", eventID)
}

type recommendationServiceImpl struct {
	db *gorm.DB
}

func NewRecommendationService(db *gorm.DB) RecommendationService {
	return &recommendationServiceImpl{db: db}
}

func (s *recommendationServiceImpl) Predict(allEvents []*models.Event, userId uint, count int) ([]uint, error) {
	userPreferences, err := s.GetUserPreferences(userId)
	if err != nil {
		log.Printf("Error getting userPreferences: %v", err)
		return nil, err
	}

	userTags := make(map[string]struct{})
	for _, tag := range userPreferences.Tags {
		userTags[tag.Name] = struct{}{}
		println(tag.Name)
	}

	recommendedEvents := getRecommendedEvents(s, allEvents, userPreferences, count)

	for _, eventID := range recommendedEvents {
		event, err := findEventByID(allEvents, eventID)
		if err != nil {
			log.Printf("Error finding event: %v", err)
			return nil, err
		}

		for _, tag := range event.Tags {
			if _, ok := userTags[tag.Name]; ok {
				log.Print(tag.Name)
			}
		}
		log.Println()
	}

	return recommendedEvents, nil
}

func (s *recommendationServiceImpl) GetUserPreferences(userId uint) (*models.UserPreference, error) {
	var userPreferences models.UserPreference
	if err := s.db.Preload("Tags").First(&userPreferences, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &userPreferences, nil
}

func (s *recommendationServiceImpl) ModifyAttendancePreference(userId uint, eventId uint, positive bool) error {
	userPreferences, err := s.GetUserPreferences(userId)
	if err != nil {
		return err
	}

	var event models.Event
	if err := s.db.Preload("Tags").First(&event, eventId).Error; err != nil {
		return err
	}

	// Track tags to be removed
	var tagsToRemove []models.Tag

	for _, tag := range event.Tags {
		if positive && !containsTag(userPreferences.Tags, tag) {
			userPreferences.Tags = append(userPreferences.Tags, *tag)
		} else {
			for i, userTag := range userPreferences.Tags {
				if userTag.Name == tag.Name {
					tagsToRemove = append(tagsToRemove, userTag)
					userPreferences.Tags = append(userPreferences.Tags[:i], userPreferences.Tags[i+1:]...)
					break
				}
			}
		}
	}

	if len(tagsToRemove) > 0 {
		if err := s.db.Model(&userPreferences).Association("Tags").Delete(tagsToRemove); err != nil {
			return err
		}
	}

	if err := s.db.Save(&userPreferences).Error; err != nil {
		return err
	}

	return nil
}

func containsTag(tag1 []models.Tag, tag2 *models.Tag) bool {
	for _, tag := range tag1 {
		if tag.Name == tag2.Name {
			return true
		}
	}
	return false
}
