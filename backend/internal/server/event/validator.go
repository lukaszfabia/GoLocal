package event

import (
	"backend/internal/app"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/functools"
	"backend/pkg/parsers"
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyForEventReq string

const _eventForm KeyForEventReq = "eventForm"
const _params KeyForEventReq = "params"
const _limit KeyForEventReq = "limit"

func (h *EventHandler) Validate(next http.Handler) http.Handler {
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

func post(w http.ResponseWriter, r *http.Request) context.Context {
	form, err := parsers.DecodeMultipartForm[forms.Event](r)
	if err != nil {
		return nil
	}

	if !functools.In(form.EventType, models.EventTypes) {
		app.InvalidDataResponse(w)
		return nil
	}

	ctx := context.WithValue(r.Context(), _eventForm, form)
	return ctx
}

func get(_ http.ResponseWriter, r *http.Request) context.Context {
	params := parsers.ParseURLQuery(r, forms.Event{}, "lon", "lat", "accuracy", "street", "street_number", "country", "city")

	limitStr := mux.Vars(r)["limit"]
	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = -1 // take all records
	}

	ctx := context.WithValue(r.Context(), _params, params)
	ctx = context.WithValue(ctx, _limit, limit)
	return ctx
}
