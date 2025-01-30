package recommendation_handler

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/recommendation"
	"log"
	"net/http"
	"strconv"
)

type RecommendationHandler struct {
	UserService           database.UserService
	EventService          database.EventService
	RecommendationService recommendation.RecommendationService
}

func (h *RecommendationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getRecommendations(w, r)
	default:
		app.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *RecommendationHandler) getRecommendations(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("User-Id")
	if userId == "" {
		log.Println("Unauthorized access")
		app.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	_, err := h.UserService.GetUser(userId)
	if err != nil {
		app.NewResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	preloads := []string{"Tags"}
	events, err := h.EventService.GetEvents(nil, 1000, preloads)
	if err != nil {
		app.NewResponse(w, http.StatusInternalServerError, "Error fetching events")
		return
	}

	log.Println("Events fetched")

	userIdUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		app.NewResponse(w, http.StatusBadRequest, "Invalid User-Id")
		return
	}

	recommendedEventIds, err := h.RecommendationService.Predict(events, uint(userIdUint), 10)
	if err != nil {
		app.NewResponse(w, http.StatusInternalServerError, "Error fetching survey")
		return
	}

	recommendedEvents := []*models.Event{}
	for _, id := range recommendedEventIds {
		event, err := h.EventService.GetEvent(id)
		if err != nil {
			log.Printf("Error fetching event with ID: %d", id)
			continue
		}
		recommendedEvents = append(recommendedEvents, event)
	}

	app.NewResponse(w, http.StatusOK, recommendedEvents)
}
