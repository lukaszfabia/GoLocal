package server

import (
	"backend/internal/database"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/functools"
	"backend/pkg/image"
	"backend/pkg/normalizer"
	"backend/pkg/parsers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get Events
// @Description Get events based on query parameters
// @Tags event
// @Accept json
// @Produce json
// @Router /api/auth/event [get]
func (s *Server) EventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getEvent(w, r)
	case http.MethodPost:
		s.createEvent(w, r)
	case http.MethodPatch:
	case http.MethodPut:
		s.updateEvent(w, r)
	case http.MethodDelete:
		s.deleteEvent(w, r)
	default:
		s.NewResponse(w, http.StatusBadRequest, "")
	}
}

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeMultipartForm[forms.Event](r)
	// Get user from ctx
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	if !functools.In(form.EventType, models.EventTypes) {
		s.InvalidFormResponse(w)
		return
	}

	info, _ := parsers.GetFileFromForm(r.MultipartForm, "image")

	url, _ := image.SaveImage[*image.EventImage](info)

	form.ImageURL = url

	// normalize tags
	form.Tags = normalizer.Normalizer(form.Tags)

	event, err := s.db.EventService().CreateEvent(*form)
	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	title := fmt.Sprintf("You've been organizer of %s", form.Title)
	body := "Check events info!"
	// dont send to requester
	ids := functools.Filter(func(e uint) bool {
		return e != user.ID
	}, form.Organizers)

	n := database.NewNotification(title, body, nil, ids)

	// already logged err
	s.db.NotificationService().SendPush(&n)

	s.NewResponse(w, http.StatusCreated, event)
}

func (s *Server) getEvent(w http.ResponseWriter, r *http.Request) {
	// /api/auth/event?=all
	// /api/auth/event?title={string}
	// /api/auth/event?id={uint}&title={string} //
	// /api/event/6?title=asas&event_type=workshop

	var event models.Event

	// map with tags and values for q
	params := parsers.ParseURLQuery(r, event, "lon", "lat", "accuracy", "street", "street_number", "country", "city")

	limitStr := mux.Vars(r)["limit"]
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = -1 // take all records
	}

	preloads := []string{
		"Location",
		"Location.Address",
		"Tags",
		"EventOrganizers",
		"Votes",
	}
	res, err := s.db.EventService().GetEvents(params, limit, preloads)

	if err != nil {
		s.NewResponse(w, http.StatusNotFound, "No such a events")
		return
	}

	s.NewResponse(w, http.StatusOK, res)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	params := parsers.ParseURLQuery(r, models.Event{}, "id")

	if _, ok := params["id"]; !ok && len(params) < 1 {
		s.InvalidFormResponse(w)
		return
	}

	id, ok := params["id"].(int)

	if !ok {
		s.InvalidFormResponse(w)
		return
	}

	event, err := s.db.EventService().DeleteEvent(id)
	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	s.NewResponse(w, http.StatusOK, event)
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {}
