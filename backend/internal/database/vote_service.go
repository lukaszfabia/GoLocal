package database

import (
	"backend/internal/models"
	"log"

	"gorm.io/gorm"
)

type VoteService interface {
	GetVotes(params map[string]any, limit int) ([]*models.Vote, error)
	// CreateVote(Vote forms.Vote) (models.Vote, error)
	// DeleteVote(id int) (models.Vote, error)
	// UpdateVote() (models.Vote, error)
}

type voteServiceImpl struct {
	db *gorm.DB
}

func NewVoteService(db *gorm.DB) VoteService {
	return &voteServiceImpl{
		db: db,
	}
}

func (s *service) VoteService() VoteService {
	return s.voteService
}

func (v *voteServiceImpl) GetVotes(params map[string]any, limit int) ([]*models.Vote, error) {
	q := v.db.Preload("Options").Preload("Event").Model(&models.Vote{})

	if limit > 0 && limit < 20 {
		q = q.Limit(limit)
	}

	//handle for event id
	if eventID, ok := params["eventID"]; ok {
		q = q.Joins("JOIN events ON events.id = votes.event_id").Where("event_id = ?", eventID)
	}

	if voteType, ok := params["voteType"]; ok {
		q = q.Where("vote_type = ?", voteType)
	}

	var votes []*models.Vote

	if err := q.Find(&votes).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return votes, nil
}
