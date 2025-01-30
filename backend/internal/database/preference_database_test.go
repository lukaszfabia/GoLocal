package database

import (
	"backend/internal/models"
	"backend/internal/recommendation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredict(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestPreferenceData()
	if err != nil {
		t.Error(err)
	}

	var user models.User
	if err := db.First(&user, "email = ?", "t@t.t").Error; err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	var events []*models.Event
	if err := db.Preload("Tags").Find(&events).Error; err != nil {
		t.Fatalf("failed to fetch events: %v", err)
	}

	service := recommendation.NewRecommendationService(db)
	recommended, err := service.Predict(events, user.ID, 2)
	if err != nil {
		t.Errorf("Predict returned error: %v", err)
	}
	t.Logf("Recommended event IDs: %v", recommended)
	if len(recommended) == 0 {
		t.Error("expected at least one recommended event, got none")
	}

	assert.Equal(t, 2, len(recommended))

	assert.Equal(t, events[0].ID, recommended[0])
	assert.Equal(t, events[2].ID, recommended[1])

	srv.Drop()
}

func TestRecommendationService_ModifyAttendancePreference(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestPreferenceData()

	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	var user models.User
	if err := db.First(&user, "email = ?", "t@t.t").Error; err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	var event models.Event
	if err := db.Joins("JOIN event_tags ON event_tags.event_id = events.id").
		Joins("JOIN tags ON tags.id = event_tags.tag_id").
		Where("tags.name = ?", "Music").
		Preload("Tags").
		First(&event).Error; err != nil {
		t.Fatalf("failed to fetch event: %v", err)
	}

	service := recommendation.NewRecommendationService(db)
	if err := service.ModifyAttendancePreference(user.ID, event.ID, false); err != nil {
		t.Errorf("ModifyAttendancePreference returned error: %v", err)
	}

	var userPreference models.UserPreference
	if err := db.Preload("Tags").First(&userPreference, "user_id = ?", user.ID).Error; err != nil {
		t.Fatalf("failed to fetch user preference: %v", err)
	}

	found := false
	for _, tag := range userPreference.Tags {
		if tag.Name == "Music" {
			found = true
			break
		}
	}
	if found {
		t.Error("expected tag to be removed from user preferences, but it's still there")
	}
}

func TestRecommendationService_NoUserPreferences(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestPreferenceData()

	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	var user = models.User{
		FirstName: "test",
		LastName:  "testsowski",
		Email:     "tt@t.t",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	service := recommendation.NewRecommendationService(db)
	_, err = service.GetUserPreferences(user.ID)

	assert.Error(t, err)
}

func TestRecommendationService_NoUserPreferencesPredict(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestPreferenceData()

	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	var user = models.User{
		FirstName: "test",
		LastName:  "testsowski",
		Email:     "tt@t.t",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	service := recommendation.NewRecommendationService(db)
	_, err = service.Predict(nil, user.ID, 2)

	assert.Error(t, err)
}
