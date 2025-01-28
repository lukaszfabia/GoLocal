package database

import (
	"backend/internal/forms"
	"backend/internal/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type VoteService interface {
	GetVotes(params map[string]any, limit int) ([]*models.Vote, error)
	Vote(params forms.VoteInVotingForm, user models.User) error
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

func (v *voteServiceImpl) Vote(params forms.VoteInVotingForm, user models.User) error {
	return v.db.Transaction(func(tx *gorm.DB) error {
		// get vote
		destVote := &models.Vote{}
		if err := v.db.First(destVote, "id = ?", params.VoteID).Error; err != nil {
			log.Println("Error fetching vote:", err)
			return err
		}

		// validate time
		if destVote.EndDate != nil && time.Now().After(*destVote.EndDate) {
			log.Println("Vote expired")
			return fmt.Errorf("vote expired")
		}

		// check if user already voted
		var existingAnswer models.VoteAnswer
		if err := v.db.First(&existingAnswer, "vote_id = ? AND user_id = ? AND deleted_at IS NULL ", params.VoteID, user.ID).Error; err == nil {
			log.Println("User already voted")
		}

		// check if user can change vote
		if destVote.VoteType == models.CannotChangeVote && existingAnswer.ID != 0 {
			log.Println("User cannot change vote")
			return fmt.Errorf("you can't change vote")
		}

		if existingAnswer.ID != 0 {
			existingAnswer.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
			if err := v.db.Save(&existingAnswer).Error; err != nil {
				log.Println("Failed to delete existing vote answer:", err)
				return fmt.Errorf("failed to delete existing vote answer")
			}
		}

		// get vote answer
		newAnswer := &models.VoteAnswer{
			VoteOptionID: uint(params.VoteOptionID),
			VoteID:       uint(params.VoteID),
			UserID:       user.ID,
		}
		if err := v.db.Create(newAnswer).Error; err != nil {
			log.Println("Failed to create vote answer:", err)
			return fmt.Errorf("failed to create vote answer")
		}

		// get vote option
		voteOption := &models.VoteOption{}
		if err := v.db.First(voteOption, "id = ?", params.VoteOptionID).Error; err != nil {
			log.Println("Failed to get vote options:", err)
			return fmt.Errorf("failed to get vote options")
		}

		// update vote option
		voteOption.VoteAnswers = append(voteOption.VoteAnswers, *newAnswer)
		if err := v.db.Save(voteOption).Error; err != nil {
			log.Println("Failed to update vote option:", err)
			return fmt.Errorf("failed to update vote option")
		}

		// update vote
		destVote.Options = append(destVote.Options, *voteOption)
		if err := v.db.Save(destVote).Error; err != nil {
			log.Println("Failed to update vote:", err)
			return fmt.Errorf("failed to update vote")
		}

		if voteOption.ParticipationStatus == models.NotInterested {
			dbInstance.recommendationService.ModifyAttendancePreference(user.ID, destVote.EventID, false)
		}

		return nil
	})
}

func (v *voteServiceImpl) GetVotes(params map[string]any, limit int) ([]*models.Vote, error) {
	q := v.db.
		Preload("Options").
		Preload("Options.VoteAnswers").
		Preload("Event").
		Model(&models.Vote{})

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
