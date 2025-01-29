package database

import (
	"context"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func timePtr(t time.Time) *time.Time {
	return &t
}

func MustStartPostgresContainer() (func(context.Context) error, error) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgis/postgis",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	database = dbName
	password = dbPwd
	username = dbUser

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return dbContainer.Terminate, err
	}

	host = dbHost
	port = dbPort.Port()

	return dbContainer.Terminate, err
}

func TestMain(m *testing.M) {
	teardown, err := MustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := New()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}

// func TestEvent(t *testing.T) {
// 	srv := New()

// 	srv.Sync()
// 	// create 3 users
// 	if err := srv.DummyService().TestUsers(); err != nil {
// 		t.Error(err)
// 	}

// 	form := forms.Event{
// 		Organizers:  []uint{1, 2},
// 		Title:       "Gaming Expo 2025",
// 		Description: "biggest gaming evnet",
// 		Image:       nil,
// 		StartDate:   timePtr(time.Now().Add(24 * time.Hour)),
// 		FinishDate:  timePtr(time.Now().Add(48 * time.Hour)),
// 		IsAdultOnly: false,
// 		EventType:   "WORKSHOP",
// 		Tags:        []string{" gaming", "e sport"},
// 		Lon:         "21.0122",
// 		Lat:         "52.2298",
// 	}

// 	expectedEvent := models.Event{
// 		Title:       "Gaming Expo 2025",
// 		Description: "biggest gaming evnet",
// 		StartDate:   timePtr(time.Now().Add(24 * time.Hour)),
// 		FinishDate:  timePtr(time.Now().Add(48 * time.Hour)),
// 		IsAdultOnly: false,
// 		EventType:   models.Workshop,
// 		Tags:        []*models.Tag{{Name: "gaming"}, {Name: "e_sport"}},
// 	}

// 	event, err := srv.EventService().CreateEvent(form)

// 	if err != nil {
// 		t.Errorf("Failed to create event: %s", err)
// 	}

// 	if event.IsAdultOnly != expectedEvent.IsAdultOnly {
// 		t.Error("Failed with adult only")
// 	}

// 	if event.ImageURL != nil {
// 		t.Errorf("Image should be nil and we have: %s", *event.ImageURL)
// 	}

// 	if event.Description != expectedEvent.Description {
// 		t.Errorf("Want %s have %s", expectedEvent.Description, event.Description)
// 	}

// 	if event.StartDate.Compare(*expectedEvent.StartDate) == 0 && event.FinishDate.Compare(*expectedEvent.FinishDate) == 0 {
// 		t.Error("Time are not the same")
// 	}

// 	if expectedEvent.Title != event.Title {
// 		t.Errorf("Want %s have %s", expectedEvent.Title, event.Title)
// 	}

// 	for id, org := range event.EventOrganizers {
// 		if org.ID != form.Organizers[id] {
// 			t.Errorf("Want %d have %d", form.Organizers[id], org.ID)
// 		}
// 	}

// 	for id, tag := range event.Tags {
// 		if tag.Name != expectedEvent.Tags[id].Name {
// 			t.Errorf("Want %s have %s", expectedEvent.Tags[id].Name, event.Title)
// 		}
// 	}

// 	if event.EventType != expectedEvent.EventType {
// 		t.Errorf("Want %s have %s", expectedEvent.EventType, event.EventType)
// 	}
// }

func TestUserService(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	user, err := srv.UserService().GetOrCreateUser(&users[0])

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(*user, users[0]) {
		t.Error("Users should be equal")
	}

}

func TestClose(t *testing.T) {
	srv := New()

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
