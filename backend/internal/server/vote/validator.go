package vote

import (
	"backend/internal/app"
	"backend/internal/forms"
	"backend/pkg/parsers"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyForVoteReq string

const _voteForm KeyForVoteReq = "voteForm"
const _params KeyForVoteReq = "params"
const _limit KeyForVoteReq = "limit"

func (h *VoteHandler) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		switch r.Method {
		case http.MethodPost:
			ctx = post(w, r)
			if ctx != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			app.InvalidDataResponse(w)
		case http.MethodGet:
			next.ServeHTTP(w, r.WithContext(get(w, r)))
		default:
			next.ServeHTTP(w, r)
		}
	})
}

// validator for voting
func post(w http.ResponseWriter, r *http.Request) context.Context {
	form, err := parsers.DecodeMultipartForm[forms.VoteInVotingForm](r)
	if err != nil {
		app.InvalidDataResponse(w)
		return nil
	}

	return context.WithValue(r.Context(), _voteForm, form)
}

// validator for getting votes with query string
func get(w http.ResponseWriter, r *http.Request) context.Context {
	println("get")

	params := make(map[string]interface{})
	query := r.URL.Query()

	if eventID := query.Get("eventID"); eventID != "" {
		params["eventID"] = eventID
	}
	if voteType := query.Get("voteType"); voteType != "" {
		params["voteType"] = voteType
	}

	limitStr := mux.Vars(r)["limit"]
	limit := -1
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			app.NewResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid limit"})
			return nil
		}
	}

	ctx := context.WithValue(r.Context(), _params, params)
	return context.WithValue(ctx, _limit, limit)
}
