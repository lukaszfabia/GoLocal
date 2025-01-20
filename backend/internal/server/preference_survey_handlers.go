package server

import (
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"

	_ "gorm.io/gorm"
)

// @Summary Get Survey
// @Description Get the preference survey
// @Tags preference
// @Accept json
// @Produce json
// @Router /api/preference/preference-survey [get]
func (s *Server) getSurvey(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		survey, err := s.db.PreferenceSurveyService().GetSurvey(1)

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

// @Summary Handle Survey
// @Description Handle the submission of a preference survey
// @Tags preference
// @Accept json
// @Produce json
// @Router /api/preference/change-preference-survey [post]
func (s *Server) handleSurvey(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var survey models.PreferenceSurvey
		if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
			s.InvalidFormResponse(w)
			return
		}
		// Store survey in database
		// ...store survey logic...
		s.NewResponse(w, http.StatusCreated, survey)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// @Summary Handle Survey Answer
// @Description Handle the submission of survey answers
// @Tags preference
// @Accept json
// @Produce json
// @Success 201 {object} []models.PreferenceSurveyAnswer
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/preference/preference-survey/answer [post]
func (s *Server) handleSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var payload struct {
			Answers []models.PreferenceSurveyAnswer `json:"answers"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Println("Error decoding JSON:", err)
			s.InvalidFormResponse(w)
			return
		}

		answers := payload.Answers

		// TODO: Check if user has already answered the survey and other logic

		for _, answer := range answers {
			s.db.PreferenceSurveyService().SaveAnswers(&answer)
		}

		s.NewResponse(w, http.StatusCreated, answers)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
