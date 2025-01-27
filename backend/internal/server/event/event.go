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
	form, _ := r.Context().Value(_eventForm).(*forms.Event)
	user, _ := r.Context().Value("user").(*models.User)

	info, _ := parsers.GetFileFromForm(r.MultipartForm, "image")

	url, _ := image.SaveImage[*image.EventImage](info)

	form.ImageURL = url

	event, err := h.EventService.CreateEvent(*form)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusBadRequest)
		return
	}

	// create notification
	n := notifications.Notification{
		Title:    fmt.Sprintf("You've been organizer of %s", event.Title),
		Body:     "Check new events info!",
		UsersIds: form.Organizers,
		Author:   user.ID,
	}

	// already logged err
	h.NotificationService.SendPush(&n)

	app.NewResponse(w, http.StatusCreated, event)
}
