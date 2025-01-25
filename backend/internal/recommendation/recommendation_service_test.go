package recommendation

// import (
// 	"testing"

// 	"backend/internal/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// // SUGGESTION(lukasz): lepiej zrobic test handlera od razu
// func TestRecommendationService_Predict(t *testing.T) {
// 	dsn := "host=localhost user=test password=test dbname=go_local_db_test port=5000 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("failed to connect database: %v", err)
// 	}

// 	// Migrate models
// 	if err := db.AutoMigrate(&models.Recommendation{}, &models.Tag{}, &models.Event{}); err != nil {
// 		t.Fatalf("failed to migrate: %v", err)
// 	}

// 	// Create mock user preferences
// 	userPreferences := models.Recommendation{
// 		Tags: []models.Tag{
// 			{Name: "Music"},
// 			{Name: "Art"},
// 		},
// 	}
// 	if err := db.Create(&userPreferences).Error; err != nil {
// 		t.Fatalf("failed to create user preferences: %v", err)
// 	}

// 	// Create mock events
// 	events := []*models.Event{
// 		{
// 			Tags: []*models.Tag{{Name: "Music"}, {Name: "Sport"}},
// 		},
// 		{
// 			Tags: []*models.Tag{{Name: "Art"}, {Name: "Food"}},
// 		},
// 		{
// 			Tags: []*models.Tag{{Name: "Music"}, {Name: "Art"}},
// 		},
// 	}
// 	for _, e := range events {
// 		if err := db.Create(e).Error; err != nil {
// 			t.Fatalf("failed to create event: %v", err)
// 		}
// 	}

// 	// Create and invoke the RecommendationService
// 	service := NewRecommendationService(db)
// 	recommended, err := service.Predict(events, userPreferences.ID)
// 	if err != nil {
// 		t.Errorf("Predict returned error: %v", err)
// 	}

// 	t.Logf("Recommended event IDs: %v", recommended)

// 	if len(recommended) == 0 {
// 		t.Error("expected at least one recommended event, got none")
// 	}
// }
