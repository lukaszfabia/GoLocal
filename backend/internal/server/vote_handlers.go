package server

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/parsers"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) VoteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getVote(w, r)
	case http.MethodPost:
		s.vote(w, r)
	default:
		s.NewResponse(w, http.StatusBadRequest, nil)
		return
	}

}

func (s *Server) vote(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.VoteForm](r)

	if err != nil {
		s.InvalidFormResponse(w)
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	if err := s.db.VoteService().Vote(*form, *user); err != nil {
		s.NewResponse(w, http.StatusBadRequest, err)
		return
	}

	s.NewResponse(w, http.StatusOK, "Successfully changed")
}

func (s *Server) getVote(w http.ResponseWriter, r *http.Request) {
	// GET api/auth/votes/?=all
	// GET api/auth/votes/?eventID=2
	// GET api/auth/votes/5?eventID=2&voteType=jakistamtyp

	var vote models.Vote

	params := parsers.ParseURLQuery(r, vote, "eventID", "voteType")

	limitStr := mux.Vars(r)["limit"]
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = 5
	}

	if res, err := s.db.VoteService().GetVotes(params, limit); err != nil {
		log.Println(err)
		s.NewResponse(w, http.StatusBadRequest, nil)
	} else {
		s.NewResponse(w, http.StatusOK, res)
	}
}
