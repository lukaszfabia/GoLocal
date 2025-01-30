package recommendation_handler

import (
	"backend/internal/forms"
	"backend/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(conds ...interface{}) (*models.User, error) {
	args := m.Called(conds)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) AddDevice(*forms.Device) error {
	return nil
}

func (m *MockUserService) GetOrCreateUser(user *models.User) (*models.User, error) {
	return nil, nil
}

func (m *MockUserService) SaveUser(new *forms.EditAccount, old *models.User) (*models.User, error) {
	return nil, nil
}

func (m *MockUserService) DeleteUser(user *models.User) error {
	return nil
}

func (m *MockUserService) VerifyUser(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (m *MockUserService) ChangePassword(newPassword string, user models.User) error {
	return nil
}

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) GetEvents(params map[string]any, limit int, preloads []string) ([]*models.Event, error) {
	args := m.Called(params, limit, preloads)
	return args.Get(0).([]*models.Event), args.Error(1)
}

func (m *MockEventService) GetEvent(eventId uint) (*models.Event, error) {
	args := m.Called(eventId)
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventService) CreateEvent(event forms.Event) (models.Event, error) {
	return models.Event{}, nil
}

func (m *MockEventService) DeleteEvent(id int) (models.Event, error) {
	return models.Event{}, nil
}

func (m *MockEventService) UpdateEvent() (models.Event, error) {
	return models.Event{}, nil
}

func (m *MockEventService) PromoteEvent(id int) (*models.Event, error) {
	return nil, nil
}

func (m *MockEventService) ReportEvent(form forms.ReportForm) error {
	return nil
}

type MockRecommendationService struct {
	mock.Mock
}

func (m *MockRecommendationService) Predict(events []*models.Event, userId uint, limit int) ([]uint, error) {
	args := m.Called(events, userId, limit)
	return args.Get(0).([]uint), args.Error(1)
}

func (m *MockRecommendationService) ModifyAttendancePreference(userId uint, eventId uint, positive bool) error {
	return nil
}

func (m *MockRecommendationService) GetUserPreferences(userId uint) (*models.UserPreference, error) {
	return nil, nil
}

func TestGetRecommendations_Success(t *testing.T) {
	mockUserService := new(MockUserService)
	mockEventService := new(MockEventService)
	mockRecommendationService := new(MockRecommendationService)

	handler := &RecommendationHandler{
		UserService:           mockUserService,
		EventService:          mockEventService,
		RecommendationService: mockRecommendationService,
	}

	user := &models.User{FirstName: "1"}
	mockUserService.On("GetUser", mock.Anything).Return(user, nil)

	events := []*models.Event{{Title: "1"}, {Title: "2"}}
	mockEventService.On("GetEvents", mock.Anything, 1000, []string{"Tags"}).Return(events, nil)
	mockRecommendationService.On("Predict", events, uint(1), 10).Return([]uint{1, 2}, nil)
	mockEventService.On("GetEvent", uint(1)).Return(&models.Event{Title: "1"}, nil)
	mockEventService.On("GetEvent", uint(2)).Return(&models.Event{Title: "2"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/recommendations", nil)
	req.Header.Set("User-Id", "1")

	w := httptest.NewRecorder()
	handler.getRecommendations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRecommendations_Unauthorized(t *testing.T) {
	handler := &RecommendationHandler{}

	req := httptest.NewRequest(http.MethodGet, "/api/recommendations", nil)
	w := httptest.NewRecorder()
	handler.getRecommendations(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetRecommendations_InternalServerError(t *testing.T) {
	mockUserService := new(MockUserService)
	handler := &RecommendationHandler{
		UserService: mockUserService,
	}

	mockUserService.On("GetUser", mock.Anything).Return(&models.User{}, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/recommendations", nil)
	req.Header.Set("User-Id", "1")

	w := httptest.NewRecorder()
	handler.getRecommendations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
