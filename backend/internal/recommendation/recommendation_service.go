package recommendation

import (
	"backend/internal/models"
	"log"

	"gorm.io/gorm"
)

type RecommendationService interface {
	Predict([]*models.Event, uint, int) ([]uint, error)
	ModifyAttendancePreference(uint, uint, bool) error
	GetUserPreferences(uint) (*models.UserPreference, error)
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
	}

	recommendedEvents := getRecommendedEvents(s, allEvents, userPreferences, count)

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

	for _, tag := range event.Tags {
		if positive && !containsTag(userPreferences.Tags, tag) {
			userPreferences.Tags = append(userPreferences.Tags, *tag)
		} else {
			for i, userTag := range userPreferences.Tags {
				if userTag.Name == tag.Name {
					userPreferences.Tags = append(userPreferences.Tags[:i], userPreferences.Tags[i+1:]...)
					break
				}
			}
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
