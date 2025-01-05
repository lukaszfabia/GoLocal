package database

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type EventService interface {
	GetEvents() ([]models.Event, error)
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

func (e *eventServiceImpl) GetEvents() ([]models.Event, error) { return nil, nil }
func (e *eventServiceImpl) CreateEvent() (models.Event, error) { return models.Event{}, nil }
func (e *eventServiceImpl) DeleteEvent() (models.Event, error) { return models.Event{}, nil }
func (e *eventServiceImpl) UpdateEvent() (models.Event, error) { return models.Event{}, nil }
