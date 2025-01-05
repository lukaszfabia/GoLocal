package services

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type RecommendationService struct {
	db *gorm.DB
}

func NewRecommendationService(db *gorm.DB) *RecommendationService {
	return &RecommendationService{db: db}
}

func (s *RecommendationService) CreateRecommendation(recommendation *models.Recommendation) error {
	return s.db.Create(recommendation).Error
}
