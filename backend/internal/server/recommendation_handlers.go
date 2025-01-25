package server

import (
	"fmt"
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

		events, err := s.db.EventService().GetEvents(map[string]any{}, 0)
		for _, event := range events {
			fmt.Println(event)
		}
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching events")
			return
		}

		userIdUint, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			s.NewResponse(w, http.StatusBadRequest, "Invalid User-Id")
			return
		}

		survey, err := s.db.RecommendationService().Predict(events, uint(userIdUint))
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching survey")
			return
		}
		s.NewResponse(w, http.StatusOK, survey)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
