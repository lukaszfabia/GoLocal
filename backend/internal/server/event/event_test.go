package event

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/notifications"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
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

func (m *MockEventService) PromoteEvent(id int) (*models.Event, error) {
	args := m.Called(id)

	if args.Get(0) != nil {
		e := args.Get(0).(*models.Event)
		e.IsPromoted = true

		return e, args.Error(1)
	}

	return nil, fmt.Errorf("event not found")
}

func (m *MockEventService) ReportEvent(form forms.ReportForm) error {
	args := m.Called(form)
	return args.Error(0)
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

func TestPost_InvalidForm(t *testing.T) {
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

func TestEventPromoValidator(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mockEventService := new(MockEventService)
	mockNotificationService := new(MockNotificationService)
	handler := &EventHandler{
		EventService:        mockEventService,
		NotificationService: mockNotificationService,
	}

	middleware := handler.ValidatePromoteEvent(mockHandler)

	t.Run("Valid PUT request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/event/123/promo", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", rr.Code)
		}
	})

	t.Run("Invalid GET request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/event/123/promo", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status BadRequest; got %v", rr.Code)
		}
	})
}

func TestEventReportValidator(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form, ok := r.Context().Value(_reportForm).(*forms.ReportForm)
		if !ok {
			t.Error("expected report form in context")
		}

		if form.ID != 123 || form.Reason != "spam" {
			t.Errorf("expected ID=123 and Reason='spam'; got ID=%v, Reason='%v'", form.ID, form.Reason)
		}

		w.WriteHeader(http.StatusOK)
	})

	// Tworzymy instancję EventHandler (jeśli jest potrzebna)
	handler := &EventHandler{}

	// Tworzymy middleware
	middleware := handler.ValidateReportEvent(mockHandler)

	t.Run("Valid POST request", func(t *testing.T) {
		jsonBody := `{"id": 123, "reason": "spam"}`
		req, err := http.NewRequest(http.MethodPost, "/event/report", strings.NewReader(jsonBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", rr.Code)
		}
	})

	t.Run("Invalid PUT request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/event/report", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status BadRequest; got %v", rr.Code)
		}
	})

	t.Run("Invalid JSON request (missing fields)", func(t *testing.T) {
		invalidJSON := `{"id": 123}`
		req, err := http.NewRequest(http.MethodPost, "/event/report", strings.NewReader(invalidJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status BadRequest; got %v", rr.Code)
		}
	})

	t.Run("Invalid JSON request (malformed JSON)", func(t *testing.T) {
		invalidJSON := `{"id": 123, "reason": "spam"`
		req, err := http.NewRequest(http.MethodPost, "/event/report", strings.NewReader(invalidJSON))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		middleware.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status BadRequest; got %v", rr.Code)
		}
	})
}

func TestPromoteEvent_Success(t *testing.T) {
	mockService := new(MockEventService)
	handler := &EventHandler{EventService: mockService}

	expectedEvent := &models.Event{Title: "Test Event"}
	expectedEvent.ID = 123

	mockService.On("PromoteEvent", 123).Return(expectedEvent, nil)

	req, err := http.NewRequest(http.MethodPost, "/events/123/promote", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	rr := httptest.NewRecorder()

	handler.PromoteEvent(rr, req)

	t.Logf("Response Body: %s", rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code)

	var response struct {
		Status int          `json:"status"`
		Data   models.Event `json:"data"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedEvent.ID, response.Data.ID)
	assert.Equal(t, expectedEvent.Title, response.Data.Title)
	assert.True(t, response.Data.IsPromoted)

	mockService.AssertCalled(t, "PromoteEvent", 123)
}

func TestPromoteEvent_InvalidID(t *testing.T) {
	mockService := new(MockEventService)
	handler := &EventHandler{EventService: mockService}

	req, err := http.NewRequest(http.MethodPost, "/events/invalid/promote", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

	rr := httptest.NewRecorder()

	handler.PromoteEvent(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var response struct {
		Status int               `json:"status"`
		Data   map[string]string `json:"data"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "invalid event ID", response.Data["message"])
}

func TestPromoteEvent_NotFound(t *testing.T) {
	mockService := new(MockEventService)
	handler := &EventHandler{EventService: mockService}

	mockService.On("PromoteEvent", 123).Return(nil, errors.New("event not found"))

	req, err := http.NewRequest(http.MethodPost, "/events/123/promote", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	rr := httptest.NewRecorder()

	handler.PromoteEvent(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	var response struct {
		Status int               `json:"status"`
		Data   map[string]string `json:"data"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "event not found", response.Data["message"])

	mockService.AssertCalled(t, "PromoteEvent", 123)
}

func TestReportEvent_Success(t *testing.T) {
	mockService := new(MockEventService)
	handler := &EventHandler{EventService: mockService}

	// Prepare the report form
	reportForm := &forms.ReportForm{
		ID:     123,
		Reason: "Inappropriate content",
	}

	mockService.On("ReportEvent", *reportForm).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/events/report", nil)
	ctx := context.WithValue(req.Context(), _reportForm, reportForm)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	handler.ReportEvent(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	mockService.AssertCalled(t, "ReportEvent", *reportForm)
}

func TestReportEvent_Failure(t *testing.T) {
	mockService := new(MockEventService)
	handler := &EventHandler{EventService: mockService}

	reportForm := &forms.ReportForm{
		ID:     123,
		Reason: "Inappropriate content",
	}

	mockService.On("ReportEvent", *reportForm).Return(errors.New("report failed"))

	req := httptest.NewRequest(http.MethodPost, "/events/report", nil)
	ctx := context.WithValue(req.Context(), _reportForm, reportForm)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	handler.ReportEvent(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	mockService.AssertCalled(t, "ReportEvent", *reportForm)
}
