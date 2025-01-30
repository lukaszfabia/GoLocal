package recommendation_handler

import (
	"backend/internal/app"
	"context"
	"net/http"
)

type KeyForRecommendationReq string

const _userId KeyForRecommendationReq = "userId"

func (h *RecommendationHandler) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		switch r.Method {
		case http.MethodGet:
			ctx = get(w, r)
			if ctx != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			app.InvalidDataResponse(w)
		default:
			next.ServeHTTP(w, r)
		}
	})
}

// validator for getting recommendations
func get(w http.ResponseWriter, r *http.Request) context.Context {
	userId := r.Header.Get("User-Id")
	if userId == "" {
		app.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return nil
	}

	return context.WithValue(r.Context(), _userId, userId)
}
