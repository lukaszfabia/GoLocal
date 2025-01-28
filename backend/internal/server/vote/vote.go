package vote

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/parsers"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type VoteHandler struct {
	VoteService database.VoteService
	UserService database.UserService
}

func (h *VoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getVotes(w, r)
	case http.MethodPost:
		h.vote(w, r)
	default:
		app.NewResponse(w, http.StatusBadRequest, nil)
		return
	}
}

func (h *VoteHandler) vote(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting vote handler")

	form, ok := r.Context().Value(_voteForm).(*forms.VoteInVotingForm)
	if !ok {
		log.Println("Error: Form data not found in context")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userId := r.Header.Get("User-Id")
	if userId == "" {
		log.Println("Unauthorized access: missing User-Id header")
		app.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}
	log.Println("User ID from header:", userId)

	user, err := h.UserService.GetUser(userId)
	if err != nil {
		log.Println("Error fetching user:", err)
		app.NewResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	log.Println("Fetched user:", user)

	newAnswer, err := h.VoteService.Vote(*form, *user)
	if err != nil {
		log.Println("Vote error:", err)
		app.NewResponse(w, http.StatusBadRequest, err)
		return
	}
	log.Println("Vote successful, new answer:", newAnswer)

	app.NewResponse(w, http.StatusOK, "Successfully changed")
	log.Println("Vote handler completed successfully")
}

func (h *VoteHandler) getVotes(w http.ResponseWriter, r *http.Request) {
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
		app.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	user, err := h.UserService.GetUser(userId)
	if err != nil {
		log.Println(err)
		message := "Internal server error"
		if err.Error() == "you can't change vote" {
			message = "You tried to change vote on a vote that doesn't allow changing votes"
		}
		app.NewResponse(w, http.StatusInternalServerError, message)
		return
	}

	votes, err := h.VoteService.GetVotes(params, limit)
	if err != nil {
		log.Println(err)
		app.NewResponse(w, http.StatusBadRequest, nil)
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

	app.NewResponse(w, http.StatusOK, transformedVotes)
}
