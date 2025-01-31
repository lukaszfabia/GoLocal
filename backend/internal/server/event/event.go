package event

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/notifications"
	"backend/pkg/image"
	"backend/pkg/parsers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	EventService        database.EventService
	NotificationService notifications.NotificationService
}

// @Summary Handle Events
// @Description Handle requests about events
// @Tags events
// @Produce json
// @Success 201 {object} models.Event
// @Success 200 {object} []models.Event
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 405 {object} map[string]string
func (h *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.post(w, r)
		return
	}
	if r.Method == "GET" {
		h.get(w, r)
		return
	} else {
		app.NewResponse(w, http.StatusMethodNotAllowed, nil)
		return
	}
}

// @Summary Get events
// @Description Get list of events with filters
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} []models.Event
// @Failure 404 {object} map[string]string
// @Router /api/auth/event/ [get]
func (h *EventHandler) get(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(_params).(map[string]any)
	limit := r.Context().Value(_limit).(int)

	res, err := h.EventService.GetEvents(params, limit, []string{
		"Location",
		"Location.Address",
		"Tags",
		"EventOrganizers",
		"Votes",
	})

	if err != nil {
		app.NewResponse(w, http.StatusNotFound, nil)
		return
	}

	app.NewResponse(w, http.StatusOK, res)
}

// @Summary Create an event
// @Description Create a new event with given data
// @Tags events
// @Accept multipart/form-data
// @Produce json
// @Param event body forms.Event true "Event data"
// @Success 201 {object} models.Event
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/event/ [post]
func (h *EventHandler) post(w http.ResponseWriter, r *http.Request) {
	form, ok := r.Context().Value(_eventForm).(*forms.Event)
	if !ok {
		log.Println("Error: Form data not found in context")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		log.Println("Error: User data not found in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	info, _ := parsers.GetFileFromForm(r.MultipartForm, "image")

	url, _ := image.SaveImage[*image.EventImage](info)

	form.ImageURL = url

	event, err := h.EventService.CreateEvent(*form)
	if err != nil {
		log.Println("Error creating event:", err)
		app.NewResponse(w, http.StatusBadRequest, nil)
		return
	}

	notification := &notifications.Notification{
		Title:    fmt.Sprintf("You've been organizer of %s", form.Title),
		Body:     "Check new events info!",
		UsersIds: form.Organizers,
		Author:   user.ID,
	}

	h.NotificationService.SendPush(notification)

	app.NewResponse(w, http.StatusCreated, event)
}

// @Summary Promote an event
// @Description Promote an existing event by ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/auth/event/{id}/promo [post]
func (h *EventHandler) PromoteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.NewResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "invalid event ID",
		})
		return
	}

	event, err := h.EventService.PromoteEvent(id)
	if err != nil {
		app.NewResponse(w, http.StatusNotFound, map[string]interface{}{
			"message": "event not found",
		})
		return
	}

	app.NewResponse(w, http.StatusOK, event)
}

// @Summary Report an event
// @Description Report an event issue
// @Tags events
// @Accept json
// @Produce json
// @Param report body forms.ReportForm true "Report details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/auth/event/report [post]
func (h *EventHandler) ReportEvent(w http.ResponseWriter, r *http.Request) {
	form, _ := r.Context().Value(_reportForm).(*forms.ReportForm)

	if err := h.EventService.ReportEvent(*form); err != nil {
		app.NewResponse(w, http.StatusBadRequest, nil)
		return
	}

	app.NewResponse(w, http.StatusCreated, nil)
}
