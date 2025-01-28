package database

import (
	"backend/internal/models"
	"log"

	"gorm.io/gorm"
)

type PreferenceSurveyService interface {
	GetSurvey(id uint) (*models.PreferenceSurvey, error)
	SaveSurvey(survey *models.PreferenceSurvey) (*models.PreferenceSurvey, error)
	DeleteSurvey(survey *models.PreferenceSurvey) error
	SaveAnswers(answer []models.PreferenceSurveyAnswer) error
}

func NewPreferenceSurveyService(db *gorm.DB) PreferenceSurveyService {
	return &preferenceSurveyServiceImpl{db: db}
}

type preferenceSurveyServiceImpl struct {
	db *gorm.DB
}

func (s *preferenceSurveyServiceImpl) GetSurvey(id uint) (*models.PreferenceSurvey, error) {
	var survey models.PreferenceSurvey
	if err := s.db.Preload("Questions.Options").First(&survey).Error; err != nil {
		log.Printf("Couldn't find survey with id %d: %v", id, err)
		return nil, err
	}
	return &survey, nil
}

func (s *preferenceSurveyServiceImpl) SaveSurvey(survey *models.PreferenceSurvey) (*models.PreferenceSurvey, error) {
	if err := s.db.Save(&survey).Error; err != nil {
		log.Printf("Couldn't save survey with id %d: %v", survey.ID, err)
		return nil, err
	}
	return survey, nil
}

func (s *preferenceSurveyServiceImpl) DeleteSurvey(survey *models.PreferenceSurvey) error {
	if err := s.db.Delete(&survey).Error; err != nil {
		log.Printf("Couldn't delete survey with id %d: %v", survey.ID, err)
		return err
	}
	return nil
}

func (s *preferenceSurveyServiceImpl) SaveAnswers(answers []models.PreferenceSurveyAnswer) error {
	for _, answer := range answers {
		if err := s.db.Create(&answer).Error; err != nil {
			return err
		}

		tags := []models.Tag{}
		for _, option := range answer.SelectedOptions {
			option.Answer = answer
			if err := s.db.First(&option.Option, option.OptionID).Error; err != nil {
				log.Printf("Couldn't find option with id %d: %v", option.OptionID, err)
				return err
			}

			if err := s.db.Save(&option).Error; err != nil {
				log.Printf("Couldn't save option with id %d: %v", option.ID, err)
				return err
			}

			tag := models.Tag{}
			if err := s.db.First(&tag, option.Option.TagID).Error; err != nil {
				log.Printf("Couldn't find tag with id %d: %v", option.Option.TagID, err)
				return err
			}

			tags = append(tags, tag)
		}

		recommendation := models.UserPreference{
			UserID: answer.UserID,
			Tags:   tags,
		}

		if err := s.db.Save(&recommendation).Error; err != nil {
			log.Printf("Couldn't save recommendation with id %d: %v", recommendation.ID, err)
			return err
		}

		if err := s.db.Model(&recommendation).Association("Tags").Replace(tags); err != nil {
			log.Printf("Couldn't save recommendation tags for recommendation with id %d: %v", recommendation.ID, err)
			return err
		}
	}
	return nil
}

func (s *service) PreferenceSurveyService() PreferenceSurveyService {
	return s.preferenceSurveyService
}
