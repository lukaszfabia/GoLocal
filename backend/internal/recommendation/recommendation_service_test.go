package recommendation

import (
	"backend/internal/models"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var allModels []any = []any{
	&models.User{},
	&models.Location{},
	&models.Coords{},
	&models.Address{},
	&models.Event{},
	&models.Tag{},
	&models.UserPreference{},
}

func CreateBaseData(t *testing.T) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "janta", "janta", "go_local_postgres_test", "5764")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	db.Exec("DROP TABLE IF EXISTS users, locations, coords, addresses, events, tags, user_preferences, event_tags, user_preferences_tags CASCADE")

	for _, model := range allModels {
		if err := db.AutoMigrate(model); err != nil {
			t.Fatalf("failed to drop table: %v", err)
		}
	}

	user := models.User{
		FirstName: "test",
		LastName:  "testsowski",
		Email:     "t@t.t",
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	tags := []*models.Tag{
		{Name: "Music"},
		{Name: "Art"},
		{Name: "Sport"},
	}

	for _, tag := range tags {
		if err := db.Create(tag).Error; err != nil {
			t.Fatalf("failed to create tag: %v", err)
		}
	}

	recommendation := &models.UserPreference{
		UserID: user.ID,
		Tags:   []models.Tag{},
	}

	var recommendation_tag1 models.Tag
	var recommendation_tag2 models.Tag
	if err := db.First(&recommendation_tag1, "name = ?", "Music").Error; err != nil {
		t.Fatalf("failed to fetch tag: %v", err)
	}
	if err := db.First(&recommendation_tag2, "name = ?", "Art").Error; err != nil {
		t.Fatalf("failed to fetch tag: %v", err)
	}

	var not_recommended_tag models.Tag
	if err := db.First(&not_recommended_tag, "name = ?", "Sport").Error; err != nil {
		t.Fatalf("failed to fetch tag: %v", err)
	}

	recommendation.Tags = append(recommendation.Tags, recommendation_tag1, recommendation_tag2)

	if err := db.Create(recommendation).Error; err != nil {
		t.Fatalf("failed to create recommendation: %v", err)
	}

	dateStr := "2021-01-01"
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}
	datePtr := &date

	longitude := 1.
	latitude := 1.
	point := fmt.Sprintf("SRID=4326;POINT(%f %f)", longitude, latitude)

	location := models.Location{
		Coords: &models.Coords{
			Longitude: longitude,
			Latitude:  latitude,
			Geom:      point,
		},
		Address: &models.Address{
			Street:         "Marszałkowska",
			StreetNumber:   "1",
			AdditionalInfo: "Warszawa",
		},
	}

	events := []*models.Event{
		{
			StartDate: datePtr,
			Tags:      []*models.Tag{&recommendation_tag1, &recommendation_tag2},
			Location:  &location,
		},
		{
			StartDate: datePtr,
			Tags:      []*models.Tag{&not_recommended_tag},
			Location:  &location,
		},
		{
			StartDate: datePtr,
			Tags:      []*models.Tag{&recommendation_tag1},
			Location:  &location,
		},
	}
	for _, e := range events {
		if err := db.Create(e).Error; err != nil {
			t.Fatalf("failed to create event: %v", err)
		}
	}

	return db, nil
}

func TestRecommendationService_Predict(t *testing.T) {
	db, err := CreateBaseData(t)
	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	var user models.User
	if err := db.First(&user, "email = ?", "t@t.t").Error; err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	var events []*models.Event
	if err := db.Preload("Tags").Find(&events).Error; err != nil {
		t.Fatalf("failed to fetch events: %v", err)
	}

	service := NewRecommendationService(db)
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
}

func TestRecommendationService_ModifyAttendancePreference(t *testing.T) {
	db, err := CreateBaseData(t)
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

	service := NewRecommendationService(db)
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
	db, err := CreateBaseData(t)
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

	service := NewRecommendationService(db)
	_, err = service.GetUserPreferences(user.ID)

	assert.Error(t, err)
}

func TestRecommendationService_NoUserPreferencesPredict(t *testing.T) {
	db, err := CreateBaseData(t)
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

	service := NewRecommendationService(db)
	_, err = service.Predict(nil, user.ID, 2)

	assert.Error(t, err)
}
