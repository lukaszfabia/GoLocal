package server

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/parsers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) VoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getVotes(w, r)
	case http.MethodPost:
		s.vote(w, r)
	default:
		s.NewResponse(w, http.StatusBadRequest, nil)
		return
	}

}

func (s *Server) vote(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.VoteInVotingForm](r)

	if err != nil {
		log.Println(err)
		s.InvalidFormResponse(w)
		return
	}

	userId := r.Header.Get("User-Id")
	if userId == "" {
		log.Println("Unauthorized access")
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	user, err := s.db.UserService().GetUser(userId)
	if err != nil {
		log.Println(err)
		s.NewResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if _, err := s.db.VoteService().Vote(*form, *user); err != nil {
		log.Println("Vote error:", err)
		s.NewResponse(w, http.StatusBadRequest, err)
		return
	}

	s.NewResponse(w, http.StatusOK, "Successfully changed")
}

func (s *Server) getVotes(w http.ResponseWriter, r *http.Request) {
	var vote models.Vote

	params := parsers.ParseURLQuery(r, vote, "eventID", "voteType")

	limitStr := mux.Vars(r)["limit"]
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = 5
	}

	userId := r.Header.Get("User-Id")
	if userId == "" {
		log.Println("Unauthorized access")
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	user, err := s.db.UserService().GetUser(userId)
	if err != nil {
		log.Println(err)
		message := "Internal server error"
		if err.Error() == "you can't change vote" {
			message = "You tried to change vote on a vote that doesn't allow changing votes"
		}
		s.NewResponse(w, http.StatusInternalServerError, message)
		return
	}

	votes, err := s.db.VoteService().GetVotes(params, limit)
	if err != nil {
		log.Println(err)
		s.NewResponse(w, http.StatusBadRequest, nil)
		return
	}

	var transformedVotes []forms.VoteForm
	for _, vote := range votes {
		var options []forms.VoteOptionForm
		for _, option := range vote.Options {
			isSelected := false
			for _, answer := range option.VoteAnswers {
				if answer.UserID == user.ID {
					isSelected = true
					break
				}
			}
			options = append(options, forms.VoteOptionForm{
				ID:         int(option.ID),
				VoteID:     int(option.VoteID),
				Text:       option.Text,
				IsSelected: isSelected,
				VotesCount: len(option.VoteAnswers),
			})
		}
		var endDate time.Time
		if vote.EndDate != nil {
			endDate = *vote.EndDate
		}

		transformedVotes = append(transformedVotes, forms.VoteForm{
			ID:       int(vote.ID),
			EventID:  int(vote.EventID),
			Text:     vote.Text,
			VoteType: string(vote.VoteType),
			EndDate:  endDate,
			Options:  options,
			Event:    vote.Event,
		})
	}

	s.NewResponse(w, http.StatusOK, transformedVotes)
}
