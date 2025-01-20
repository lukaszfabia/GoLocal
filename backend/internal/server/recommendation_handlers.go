package server

import (
	"net/http"
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
		events, err := s.db.EventService().GetEvents(map[string]any{}, 0)
		survey, err := s.db.RecommendationService().Predict(events, "1")

		if err != nil {
			s.NewResponse(w, http.StatusInternalServerError, "Error fetching survey")
			return
		}

		// TODO: check if user did not fill out survey

		s.NewResponse(w, http.StatusOK, survey)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
