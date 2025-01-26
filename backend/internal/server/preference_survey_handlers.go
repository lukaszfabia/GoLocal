package server

import (
	"backend/internal/database"
	"backend/internal/forms"
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
			Answers []forms.PreferenceSurveyAnswer `json:"answers"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Println("Error decoding JSON:", err)
			s.InvalidFormResponse(w)
			return
		}

		answers := payload.Answers

		// Map DTO to Model
		var modelAnswers []models.PreferenceSurveyAnswer
		for _, answer := range answers {
			modelAnswer := models.PreferenceSurveyAnswer{
				SurveyID:   uint(answer.PreferenceSurveyID),
				QuestionID: uint(answer.QuestionID),
				UserID:     uint(answer.UserID),
			}
			for _, optionID := range answer.Options {
				modelAnswer.SelectedOptions = append(modelAnswer.SelectedOptions, models.PreferenceSurveyAnswerOption{
					OptionID: uint(optionID),
					Answer:   modelAnswer,
				})
			}
			modelAnswers = append(modelAnswers, modelAnswer)
		}

		// Save answers to database
		if err := s.db.PreferenceSurveyService().SaveAnswers(modelAnswers); err != nil {
			log.Println("Error saving answers:", err)
			s.NewResponse(w, http.StatusInternalServerError, "Error saving answers")
			return
		}

		log.Println("Received answers:", answers)

		title := "Hey, you just filled out the survey!"
		body := "Check info!"

		var userId = modelAnswers[0].UserID

		ids := []uint{userId}

		n := database.NewNotification(title, body, nil, ids)

		s.db.NotificationService().SendPush(&n)

		log.Println("Sent push notification")

		s.NewResponse(w, http.StatusCreated, answers)
	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
