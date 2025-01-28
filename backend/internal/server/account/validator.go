package account

import (
	"backend/internal/app"
	"backend/internal/forms"
	"backend/pkg/parsers"
	"context"
	"net/http"
)

type userKey string

const (
	_userForm userKey = "userForm"
)

func (h *AccountHandler) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		switch r.Method {
		case http.MethodPut:
			if ctx = put(w, r); ctx != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			app.InvalidDataResponse(w)
			return
		default:
			next.ServeHTTP(w, r)
		}
	})
}

func put(w http.ResponseWriter, r *http.Request) context.Context {
	form, err := parsers.DecodeMultipartForm[forms.EditAccount](r)
	if err != nil {
		app.InvalidDataResponse(w)
		return nil
	}

	return context.WithValue(r.Context(), _userForm, form)
}
