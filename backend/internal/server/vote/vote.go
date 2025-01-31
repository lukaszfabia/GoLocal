package vote

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/forms"
	"log"
	"net/http"
	"time"
)

type VoteHandler struct {
	VoteService database.VoteService
	UserService database.UserService
}

// @Summary Handle votes
// @Description Handle voting requests
// @Tags votes
// @Accept json
// @Produce json
// @Success 200 {object} []forms.VoteForm
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 405 {object} map[string]string
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

// @Summary Cast a vote
// @Description Submit a user's vote
// @Tags votes
// @Accept json
// @Produce json
// @Param vote body forms.VoteInVotingForm true "Vote data"
// @Param User-Id header string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/vote/ [post]
func (h *VoteHandler) vote(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting vote handler")

	form, ok := r.Context().Value(_voteForm).(*forms.VoteInVotingForm)
	if !ok {
		log.Println("Error: Form data not found in context")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Println("Form data retrieved from context:", form)

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

// @Summary Get votes
// @Description Retrieve votes with filters
// @Tags votes
// @Accept json
// @Produce json
// @Param limit path int true "Number of results"
// @Param User-Id header string true "User ID"
// @Success 200 {object} []forms.VoteForm
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/vote/{limit} [get]
func (h *VoteHandler) getVotes(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value(_params).(map[string]interface{})
	if !ok {
		log.Println("Error: Params not found in context")
		app.NewResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	limit, ok := r.Context().Value(_limit).(int)
	if !ok {
		log.Println("Error: Limit not found in context")
		app.NewResponse(w, http.StatusBadRequest, "Invalid request")
		return
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
