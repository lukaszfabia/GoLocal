package server

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	form, err := pkg.DecodeMultipartForm[forms.Event](r)

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	if !pkg.In(form.EventType, models.EventTypes) {
		s.InvalidFormResponse(w)
		return
	}

	info, _ := pkg.GetFileFromForm(r.MultipartForm, "image")

	url, _ := pkg.SaveImage[*pkg.EventImage](info)

	form.ImageURL = url

	if event, err := s.db.EventService().CreateEvent(*form); err != nil {
		s.InvalidFormResponse(w)
		return
	} else {
		s.NewResponse(w, http.StatusCreated, event)
		return
	}

}

func (s *Server) getEvent(w http.ResponseWriter, r *http.Request) {
	// /api/auth/event?=all
	// /api/auth/event?title={string}
	// /api/auth/event?id={uint}&title={string} //
	// /api/event/6?title=asas&event_type=workshop

	var event models.Event

	// map with tags and values for q
	params := pkg.ParseURLQuery(r, event, "lon", "lat", "accuracy", "street", "street_number", "country", "city")

	limitStr := mux.Vars(r)["limit"]
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = -1 // take all records
	}

	res, err := s.db.EventService().GetEvents(params, limit)

	if err != nil {
		s.NewResponse(w, http.StatusNotFound, "No such a events")
		return
	}

	s.NewResponse(w, http.StatusOK, res)

}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	params := pkg.ParseURLQuery(r, models.Event{}, "id")

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
