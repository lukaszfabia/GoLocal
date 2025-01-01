package server

import (
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

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

func (s *Server) handleSurveyAnswer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var answer models.PreferenceSurveyAnswer
		if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
			s.InvalidFormResponse(w)
			return
		}

		// TODO: Check if user has already answered the survey and other logic

		s.db.PreferenceSurveyService().SaveAnswers(&answer)

		s.NewResponse(w, http.StatusCreated, answer)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
