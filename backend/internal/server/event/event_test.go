package event

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/notifications"
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) GetEvents(map[string]any, int, []string) ([]*models.Event, error) {
	return nil, nil
}

func (m *MockEventService) CreateEvent(form forms.Event) (models.Event, error) {
	m.Called(form)

	return models.Event{
		Title:       form.Title,
		Description: form.Description,
		EventType:   models.EventType(form.EventType),
	}, nil
}

func (m *MockEventService) DeleteEvent(id int) (models.Event, error) {
	return models.Event{}, nil
}

func (m *MockEventService) UpdateEvent() (models.Event, error) {
	return models.Event{}, nil
}

func (m *MockEventService) GetEvent(eventId uint) (*models.Event, error) {
	return nil, nil
}

// MockNotificationService is a mock implementation of notifications.NotificationService
type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) SendPush(notification *notifications.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}

func (m *MockNotificationService) SetClient(client notifications.MessagingClient) {
	m.Called(client)
}

func TestPost_ValidForm(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("title", "Sample Event")
	writer.WriteField("description", "A great event")
	writer.WriteField("isAdultOnly", "true")
	writer.WriteField("eventType", "WORKSHOP")
	writer.WriteField("lon", "19.4326")
	writer.WriteField("lat", "52.2122")
	writer.WriteField("tags", "music")
	writer.WriteField("tags", "live")
	writer.WriteField("organizers", "1")
	writer.WriteField("organizers", "2")

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	resultCtx := post(w, req)

	assert.NotNil(t, resultCtx, "Expected context to be returned")

	form, ok := resultCtx.Value(_eventForm).(*forms.Event)
	assert.True(t, ok, "Expected form to be in context")

	assert.Equal(t, "Sample Event", form.Title, "Expected Title to match")
	assert.Equal(t, "A great event", form.Description, "Expected Description to match")
	assert.Equal(t, true, form.IsAdultOnly, "Expected IsAdultOnly to match")
	assert.Equal(t, "WORKSHOP", form.EventType, "Expected EventType to match")
	assert.Equal(t, "19.4326", form.Lon, "Expected Lon to match")
	assert.Equal(t, "52.2122", form.Lat, "Expected Lat to match")
	assert.Equal(t, []uint{1, 2}, form.Organizers, "Expected Organizers to match")
	assert.Equal(t, []string{"music", "live"}, form.Tags, "Expected Tags to match")
}

func TestPost_InvalidEventType(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("title", "Sample Event")
	writer.WriteField("description", "A great event")
	writer.WriteField("isAdultOnly", "true")
	writer.WriteField("eventType", "QATESTING")
	writer.WriteField("lon", "19.4326")
	writer.WriteField("lat", "52.2122")
	writer.WriteField("tags", "music")
	writer.WriteField("tags", "live")
	writer.WriteField("organizers", "1")
	writer.WriteField("organizers", "2")

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	resultCtx := post(w, req)

	assert.Nil(t, resultCtx, "Expected context to be returned")
}

func TestPost_InValidForm(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("title", "Sample Event")
	writer.WriteField("description", "A great event")
	writer.WriteField("info", "This should't exist")

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	resultCtx := post(w, req)

	assert.Nil(t, resultCtx, "Expected context to be returned")

}

func TestPostEvent_Success(t *testing.T) {
	mockEventService := new(MockEventService)
	mockNotificationService := new(MockNotificationService)
	handler := &EventHandler{
		EventService:        mockEventService,
		NotificationService: mockNotificationService,
	}

	form := &forms.Event{
		Title:       "Test Event",
		Organizers:  []uint{0, 2},
		EventType:   "WORKSHOP",
		Description: "Test Description",
	}

	event := models.Event{
		Title:       form.Title,
		Description: form.Description,
		EventType:   models.EventType(form.EventType),
	}

	user := &models.User{
		FirstName: "test",
		LastName:  "testsowski",
		Email:     "test@example.com",
	}

	fmt.Println("Form in test:", *form)

	mockEventService.On("CreateEvent", mock.AnythingOfType("forms.Event")).Return(event, nil)

	notification := &notifications.Notification{
		Title:    fmt.Sprintf("You've been organizer of %s", form.Title),
		Body:     "Check new events info!",
		UsersIds: form.Organizers,
		Author:   user.ID,
	}
	mockNotificationService.On("SendPush", notification).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", nil)
	ctx := context.WithValue(req.Context(), _eventForm, form)
	ctx = context.WithValue(ctx, "user", user)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	handler.post(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	mockEventService.AssertExpectations(t)
	mockNotificationService.AssertExpectations(t)
}

func TestValidate_NoForm(t *testing.T) {
	mockEventService := new(MockEventService)
	mockNotificationService := new(MockNotificationService)
	handler := &EventHandler{
		EventService:        mockEventService,
		NotificationService: mockNotificationService,
	}

	user := &models.User{
		FirstName: "test",
		LastName:  "testsowski",
		Email:     "test@example.com",
	}

	mockEventService.On("CreateEvent", mock.AnythingOfType("forms.Event")).Return(nil, nil).Maybe()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", nil)
	ctx := context.WithValue(req.Context(), "user", user)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	handler.post(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	mockEventService.AssertExpectations(t)
	mockNotificationService.AssertExpectations(t)
}

func TestValidate_Unauthorized(t *testing.T) {
	mockEventService := new(MockEventService)
	mockNotificationService := new(MockNotificationService)
	handler := &EventHandler{
		EventService:        mockEventService,
		NotificationService: mockNotificationService,
	}

	form := &forms.Event{
		Title:       "Test Event",
		Organizers:  []uint{0, 2},
		EventType:   "WORKSHOP",
		Description: "Test Description",
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/event", nil)
	ctx := context.WithValue(req.Context(), _eventForm, form)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	handler.post(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

}
