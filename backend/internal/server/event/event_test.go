package event_test

import (
	"backend/internal/app"
	"backend/internal/server/event"
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func getMethodMock(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(map[string]any)
	limit := r.Context().Value("limit").(int)

	app.NewResponse(w, http.StatusOK, map[string]any{
		"params": params,
		"limit":  limit,
	})
}

func TestEventHandler_Validate(t *testing.T) {
	handler := &event.EventHandler{}

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedCtx    map[string]any
		expectedLimit  int
		params         map[string]any
	}{
		{
			name:           "Valid request",
			url:            "/events?lon=0&lat=0&limit=10",
			expectedStatus: http.StatusOK,
			expectedCtx:    map[string]any{"lon": 0, "lat": 0},
			expectedLimit:  10,
		},
		{
			name:           "Invalid limit",
			url:            "/events?lon=0&lat=0&limit=5",
			expectedStatus: http.StatusBadRequest,
			expectedCtx:    map[string]any{"lon": 0, "lat": 0},
			expectedLimit:  5,
		},
		{
			name:           "Invalid lon/lat",
			url:            "/events?lon=1&lat=1&limit=10",
			expectedStatus: http.StatusBadRequest,
			expectedCtx:    map[string]any{"lon": 1, "lat": 1},
			expectedLimit:  10,
		},
		{
			name:           "Missing lon/lat",
			url:            "/events?limit=10",
			expectedStatus: http.StatusBadRequest,
			expectedCtx:    nil,
			expectedLimit:  10,
		},
		{
			name:           "Missing limit",
			url:            "/events?lon=0&lat=0",
			expectedStatus: http.StatusBadRequest,
			expectedCtx:    map[string]any{"lon": 0, "lat": 0},
			expectedLimit:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			req = mux.SetURLVars(req, map[string]string{
				"limit": "10",
			})

			ctx := context.WithValue(req.Context(), "params", tt.params)
			ctx = context.WithValue(ctx, "limit", tt.expectedLimit)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handlerWithValidation := handler.Validate(http.HandlerFunc(getMethodMock))

			handlerWithValidation.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				respParams := ctx.Value("params").(map[string]any)
				respLimit := ctx.Value("limit").(int)

				if !reflect.DeepEqual(respParams, tt.expectedCtx) {
					t.Errorf("Expected context params %v, got %v", tt.expectedCtx, respParams)
				}

				if respLimit != tt.expectedLimit {
					t.Errorf("Expected context limit %d, got %d", tt.expectedLimit, respLimit)
				}
			}
		})
	}
}
