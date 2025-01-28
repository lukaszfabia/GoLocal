package server

import (
	"backend/internal/models"
	"log"
	"net/http"
	"strconv"
)

// @Summary Get Recommendations
// @Description Get recommendations based on user preferences
// @Tags recommendation
// @Accept json
// @Produce json
// @Router /api/recommendation/recommendation [get]
func (s *Server) getRecommendations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userId := r.Header.Get("User-Id")
		if userId == "" {
			log.Println("Unauthorized access")
			s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
			return
		}

		_, err := s.db.UserService().GetUser(userId)
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		preloads := []string{
			"Tags",
		}
		events, err := s.db.EventService().GetEvents(nil, 1000, preloads)
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching events")
			return
		}

		log.Println("Events fetched")

		userIdUint, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			s.NewResponse(w, http.StatusBadRequest, "Invalid User-Id")
			return
		}

		recommendedEventIds, err := s.db.RecommendationService().Predict(events, uint(userIdUint), 10)
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching survey")
			return
		}

		recommendedEvents := []*models.Event{}
		for _, id := range recommendedEventIds {
			event, err := s.db.EventService().GetEvent(id)
			if err != nil {
				log.Printf("Error fetching event with ID: %d", id)
				continue
			}
			recommendedEvents = append(recommendedEvents, event)
		}

		s.NewResponse(w, http.StatusOK, recommendedEvents)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
