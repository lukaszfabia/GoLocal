package database

import (
	"backend/internal/models"
	"log"

	"gorm.io/gorm"
)

type TokenService interface {
	// Creates user if it does not exists in db
	// Returns occured error
	IsTokenBlacklisted(token string) bool
	SetAsBlacklisted(token string) error
}

type tokenServiceImpl struct {
	db *gorm.DB
}

func NewTokenService(db *gorm.DB) TokenService {
	return &tokenServiceImpl{db: db}
}

func (t *tokenServiceImpl) SetAsBlacklisted(token string) error {
	if err := t.db.Create(&models.BlacklistedTokens{Token: token}).Error; err != nil {
		log.Printf("Can not blacklist token %s\n", token)
		return err
	}

	return nil
}

func (t *tokenServiceImpl) IsTokenBlacklisted(token string) bool {
	return t.db.First(&models.BlacklistedTokens{}, "token = ?", token).Error == nil
}

func (s *service) TokenService() TokenService {
	return s.tokenService
}
