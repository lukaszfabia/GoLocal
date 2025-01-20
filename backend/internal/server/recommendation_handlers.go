package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/recommendation/"), "/")
		if len(pathParts) < 1 {
			s.NewResponse(w, http.StatusBadRequest, "Missing user ID")
			return
		}
		userID64, err := strconv.ParseUint(pathParts[0], 10, 32)
		if err != nil {
			s.NewResponse(w, http.StatusBadRequest, "Invalid user ID")
			return
		}
		userID := uint(userID64)

		events, err := s.db.EventService().GetEvents(map[string]any{}, 0)
		for _, event := range events {
			fmt.Println(event)
		}
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching events")
			return
		}
		survey, err := s.db.RecommendationService().Predict(events, userID)
		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching survey")
			return
		}
		s.NewResponse(w, http.StatusOK, survey)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
