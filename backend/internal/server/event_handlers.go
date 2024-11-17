package server

import (
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

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {}

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

	res, err := s.db.EventService().FilterEvents(params, limit)

	if err != nil {
		s.NewResponse(w, http.StatusNotFound, "No such a events")
		return
	}

	s.NewResponse(w, http.StatusOK, res)
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {}
