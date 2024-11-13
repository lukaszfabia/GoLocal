package server

import "net/http"

func (s *Server) EventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEvent(w, r)
	case http.MethodPost:
		createEvent(w, r)
	case http.MethodPatch:
	case http.MethodPut:
		updateEvent(w, r)
	case http.MethodDelete:
		deleteEvent(w, r)
	default:
		s.NewResponse(w, http.StatusBadRequest, "")
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {}

func getEvent(w http.ResponseWriter, r *http.Request) {}

func deleteEvent(w http.ResponseWriter, r *http.Request) {}

func updateEvent(w http.ResponseWriter, r *http.Request) {}
