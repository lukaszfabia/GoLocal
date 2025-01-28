package event

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/notifications"
	"fmt"
	"log"
	"net/http"
)

type EventHandler struct {
	EventService        database.EventService
	NotificationService notifications.NotificationService
}

// @Summary Get Events
// @Description Get events based on query parameters
// @Tags event
// @Accept json
// @Produce json
// @Router /api/auth/event [get]
func (h *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.post(w, r)
	}
	if r.Method == "GET" {
		h.get(w, r)
	}
}

func (h *EventHandler) get(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(_params).(map[string]any)
	limit := r.Context().Value(_limit).(int)

	log.Println(params)
	log.Println(limit)

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
