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
	SaveAnswers(answer *models.PreferenceSurveyAnswer) error
}

func NewPreferenceSurveyService(db *gorm.DB) PreferenceSurveyService {
	return &preferenceSurveyServiceImpl{db: db}
}

type preferenceSurveyServiceImpl struct {
	db *gorm.DB
}

func (s *preferenceSurveyServiceImpl) GetSurvey(id uint) (*models.PreferenceSurvey, error) {
	var survey models.PreferenceSurvey
	if err := s.db.Preload("Questions.Options").First(&survey, id).Error; err != nil {
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

func (s *preferenceSurveyServiceImpl) SaveAnswers(answer *models.PreferenceSurveyAnswer) error {
	if err := s.db.Save(&answer).Error; err != nil {
		log.Printf("Couldn't save answer with id %d: %v", answer.ID, err)
		return err
	}
	return nil
}

func (s *service) PreferenceSurveyService() PreferenceSurveyService {
	return s.preferenceSurveyService
}
