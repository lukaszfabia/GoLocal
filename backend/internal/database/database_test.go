package database

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/internal/recommendation"
	"backend/pkg/parsers"
	"context"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func timePtr(t time.Time) *time.Time {
	return &t
}

func stringPtr(s string) *string {
	return &s
}

func timeApproxEqual(t1, t2 *time.Time) bool {
	const epsilon = time.Second
	return t1.Sub(*t2) < epsilon && t2.Sub(*t1) < epsilon
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

func TestEventServiceCreate(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	sDate := timePtr(time.Now().Add(24 * time.Hour))
	fDate := timePtr(time.Now().Add(48 * time.Hour))

	form := forms.Event{
		Organizers:  []uint{users[0].ID, users[1].ID},
		Title:       "Gaming Expo 2025",
		Description: "biggest gaming evnet",
		StartDate:   sDate,
		FinishDate:  fDate,
		IsAdultOnly: false,
		EventType:   "WORKSHOP",
		Tags:        []string{" gaming", "e sport"},
		Lon:         "17.049275412738254",
		Lat:         "51.12022845",
	}

	expectedEvent := models.Event{
		Title:       "Gaming Expo 2025",
		Description: "biggest gaming evnet",
		StartDate:   sDate,
		FinishDate:  fDate,
		IsAdultOnly: false,
		EventType:   models.Workshop,
		Tags:        []*models.Tag{{Name: "gaming"}, {Name: "e_sport"}},
	}

	event, err := srv.EventService().CreateEvent(form)

	if err != nil {
		t.Errorf("Failed to create event: %s", err)
	}

	if event.IsAdultOnly != expectedEvent.IsAdultOnly {
		t.Errorf("Failed with adult only want %v have %v", expectedEvent.IsAdultOnly, event.IsAdultOnly)
	}

	if len(*event.ImageURL) != 0 {
		t.Errorf("Image should be nil and we have: %s", *event.ImageURL)
	}

	if event.Description != expectedEvent.Description {
		t.Errorf("Want %s have %s", expectedEvent.Description, event.Description)
	}

	if !timeApproxEqual(event.StartDate, expectedEvent.StartDate) {
		t.Errorf("StartDate mismatch. Expected: %v, got: %v", expectedEvent.StartDate, event.StartDate)
	}

	if !timeApproxEqual(event.FinishDate, expectedEvent.FinishDate) {
		t.Errorf("FinishDate mismatch. Expected: %v, got: %v", expectedEvent.FinishDate, event.FinishDate)
	}

	if expectedEvent.Title != event.Title {
		t.Errorf("Want %s have %s", expectedEvent.Title, event.Title)
	}

	for id, org := range event.EventOrganizers {
		if org.ID != form.Organizers[id] {
			t.Errorf("Want %d have %d", form.Organizers[id], org.ID)
		}
	}

	for id, tag := range event.Tags {
		if tag.Name != expectedEvent.Tags[id].Name {
			t.Errorf("Want %s have %s", expectedEvent.Tags[id].Name, event.Title)
		}
	}

	if event.EventType != expectedEvent.EventType {
		t.Errorf("Want %s have %s", expectedEvent.EventType, event.EventType)
	}

	srv.Drop()
}

func TestEventServiceCreate_FailureWithOrganizers(t *testing.T) {
	srv := New()

	srv.Sync()

	form := forms.Event{
		Title:       "Gaming Expo 2025",
		Description: "biggest gaming evnet",
		StartDate:   timePtr(time.Now().Add(24 * time.Hour)),
		FinishDate:  timePtr(time.Now().Add(48 * time.Hour)),
		IsAdultOnly: false,
		EventType:   "WORKSHOP",
		Tags:        []string{" gaming", "e sport"},
		Lon:         "17.049275412738254",
		Lat:         "51.12022845",
	}

	_, err := srv.EventService().CreateEvent(form)

	if err == nil {
		t.Error("Shoud be failed no organizers")
	}

	form1 := forms.Event{
		Organizers:  []uint{7},
		Title:       "Gaming Expo 2025",
		Description: "biggest gaming evnet",
		StartDate:   timePtr(time.Now().Add(24 * time.Hour)),
		FinishDate:  timePtr(time.Now().Add(48 * time.Hour)),
		IsAdultOnly: false,
		EventType:   "WORKSHOP",
		Tags:        []string{" gaming", "e sport"},
		Lon:         "17.049275412738254",
		Lat:         "51.12022845",
	}

	_, err = srv.EventService().CreateEvent(form1)

	if err == nil {
		t.Error("Orgzanier should not exists")
	}

	srv.Drop()
}

func TestEventServiceCreate_FailureLocation(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	form := forms.Event{
		Organizers:  []uint{users[0].ID, users[1].ID},
		Title:       "Gaming Expo 2025",
		Description: "biggest gaming evnet",
		StartDate:   timePtr(time.Now().Add(24 * time.Hour)),
		FinishDate:  timePtr(time.Now().Add(48 * time.Hour)),
		IsAdultOnly: false,
		EventType:   "WORKSHOP",
		Tags:        []string{" gaming", "e sport"},
	}

	_, err = srv.EventService().CreateEvent(form)

	if err == nil {
		t.Errorf("Should be failed: %s", err)
	}

	srv.Drop()
}

func TestUserServiceGetOrCreateUser_Succes(t *testing.T) {
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

	srv.Drop()

}

func TestUserServiceGetUser_Success(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	user, err := srv.UserService().GetUser("first_name = ? AND last_name = ?", users[1].FirstName, users[1].LastName)

	if err != nil {
		t.Error(err)
	}

	user.CreatedAt = user.CreatedAt.Truncate(time.Microsecond)
	users[1].CreatedAt = users[1].CreatedAt.Truncate(time.Microsecond)

	user.UpdatedAt = user.UpdatedAt.Truncate(time.Microsecond)
	users[1].UpdatedAt = users[1].UpdatedAt.Truncate(time.Microsecond)

	if (user.Bio == nil && users[1].Bio != nil) || (user.Bio != nil && users[1].Bio == nil) {
		t.Error("Bio field mismatch")
	} else if user.Bio != nil && users[1].Bio != nil && *user.Bio != *users[1].Bio {
		t.Error("Bio values are different")
	}

	if !(user.FirstName == users[1].FirstName && user.LastName == users[1].LastName) {
		t.Errorf("Not the same want %s %s have %s %s", users[1].FirstName, users[1].LastName, user.FirstName, user.LastName)
	}

	srv.Drop()

}

func TestUserServiceChangePassword_Sucess(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	myuser := users[0]

	err = srv.UserService().ChangePassword("BardzoTajnyP@ssw0rd!2d", myuser)

	if err != nil {
		t.Errorf("Something went wrong: %v", err)
	}

	if user, err := srv.UserService().GetUser("email = ?", myuser.Email); err == nil {
		if len(*user.Password) == 0 {
			t.Error("No password")
		}
	} else {
		t.Errorf("Failed to find user %s : %v", myuser.Email, err)
	}

	srv.Drop()
}

func TestUserServiceChangePassword_Failure(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	myuser := users[0]

	// weak password
	err = srv.UserService().ChangePassword("weak", myuser)

	if err == nil {
		t.Errorf("Something went wrong: %v", err)
	}

	srv.Drop()
}

func TestUserServiceGetUser_Failure(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	_, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	_, err = srv.UserService().GetUser("email = ?", "peter.fun@pwr.edu.pwr.pl.com.org.io.rus.twoj.stary")

	if err == nil {
		t.Errorf("User does not exists")
	}

	srv.Drop()

}

func TestUserServiceSaveUser_Success(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	form := &forms.EditAccount{
		Birthday:  "1970-01-01",
		Password:  "P@ssword2003",
		Email:     "joe.doe@example.com",
		FirstName: "Joe",
		LastName:  "Doeski",
		Bio:       "Przemoc domowa to nie problem, to rozwiązanie",
	}

	user, err := srv.UserService().SaveUser(form, &users[0])

	if err != nil {
		t.Error(err)
	}

	if !user.Birthday.Equal(parsers.ParseDate(form.Birthday)) {
		t.Errorf("Expected Birthday: %v, but got: %v", form.Birthday, user.Birthday)
	}

	if user.LastName != form.LastName {
		t.Errorf("Expected LastName: %s, but got: %s", form.LastName, user.LastName)
	}

	if *user.Bio != form.Bio {
		t.Errorf("Expected Bio: %v, but got: %v", form.Bio, *user.Bio)
	}

	if !(strings.HasPrefix(*user.Password, "$2a$") ||
		strings.HasPrefix(*user.Password, "$2b$")) {
		t.Errorf("Something went wrong during hashing have %s", *user.Password)
	}

	srv.Drop()
}

func TestUserServiceAddDevice_Success(t *testing.T) {
	srv := New()

	srv.Sync()
	// create 3 users
	users, err := srv.DummyService().TestUsers()
	if err != nil {
		t.Error(err)
	}

	device := &forms.Device{
		UserID:    int(users[0].ID),
		Token:     "token",
		OSVersion: stringPtr("archlinux"),
		Platform:  stringPtr("linux"),
	}

	if err := srv.UserService().AddDevice(device); err != nil {
		t.Errorf("Failed to add device to user %v", err)
	}

	user, err := srv.UserService().GetUser("email = ?", users[0].Email)

	if err != nil {
		t.Errorf("Failed to find user %v", err)
	}

	deviceFromDB := *user.Devices[0]

	if *deviceFromDB.OSVersion != *device.OSVersion {
		t.Error()
	}

	if *deviceFromDB.Platform != *device.Platform {
		t.Error()
	}

	if deviceFromDB.Token != device.Token {
		t.Error()
	}

	srv.Drop()
}

func TestUserServiceAddDevice_Failure(t *testing.T) {
	srv := New()

	srv.Sync()

	device := &forms.Device{
		UserID:    78,
		Token:     "token",
		OSVersion: stringPtr("archlinux"),
		Platform:  stringPtr("linux"),
	}

	if err := srv.UserService().AddDevice(device); err == nil {
		t.Errorf("Should failed, user does not exists")
	}

	srv.Drop()
}

func TestRecommendationService_Predict(t *testing.T) {
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

	srv.Drop()
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

	srv.Drop()
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

	srv.Drop()
}

func TestVote_Vote(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestVoteData()
	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	votingService := NewVoteService(db)

	user, err := srv.UserService().GetUser("email = ?", "t@t.t")
	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	event, err := srv.EventService().GetEvent(1)
	if err != nil {
		t.Fatalf("failed to fetch event: %v", err)
	}

	params := map[string]interface{}{
		"eventID": event.ID,
	}

	votes, err := votingService.GetVotes(params, 1)

	vote := votes[0]

	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	voteForm := forms.VoteInVotingForm{
		VoteID:       int(vote.ID),
		VoteOptionID: int(vote.Options[0].ID),
	}

	voteAnswer, err := votingService.Vote(voteForm, *user)
	assert.NoError(t, err)

	assert.Equal(t, voteAnswer.UserID, user.ID)

	srv.Drop()
}

func TestVote_VoteExpired(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestVoteData()
	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	votingService := NewVoteService(db)

	user, err := srv.UserService().GetUser("email = ?", "t@t.t")
	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	event, err := srv.EventService().GetEvent(1)
	if err != nil {
		t.Fatalf("failed to fetch event: %v", err)
	}

	params := map[string]interface{}{
		"eventID": event.ID,
	}

	votes, err := votingService.GetVotes(params, 1)

	vote := votes[0]

	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	voteForm := forms.VoteInVotingForm{
		VoteID:       int(vote.ID),
		VoteOptionID: int(vote.Options[0].ID),
	}

	endDate := time.Now().Add(-time.Hour)
	vote.EndDate = timePtr(endDate)

	if err := db.Save(&vote).Error; err != nil {
		t.Fatalf("failed to save vote: %v", err)
	}

	voteForm.VoteOptionID = int(vote.Options[1].ID)

	_, err = votingService.Vote(voteForm, *user)
	assert.Error(t, err)

	srv.Drop()
}

func TestVote_VoteEventExpired(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestVoteData()
	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	votingService := NewVoteService(db)

	user, err := srv.UserService().GetUser("email = ?", "t@t.t")
	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	event, err := srv.EventService().GetEvent(1)
	if err != nil {
		t.Fatalf("failed to fetch event: %v", err)
	}

	params := map[string]interface{}{
		"eventID": event.ID,
	}

	votes, err := votingService.GetVotes(params, 1)

	vote := votes[0]

	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	voteForm := forms.VoteInVotingForm{
		VoteID:       int(vote.ID),
		VoteOptionID: int(vote.Options[0].ID),
	}

	endDate := time.Now().Add(-time.Hour)
	event.StartDate = timePtr(endDate)

	if err := db.Save(&event).Error; err != nil {
		t.Fatalf("failed to save event: %v", err)
	}

	voteForm.VoteOptionID = int(vote.Options[1].ID)

	_, err = votingService.Vote(voteForm, *user)
	assert.Error(t, err)

	srv.Drop()
}

func TestVote_VoteChangeNotAllowed(t *testing.T) {
	srv := New()

	srv.Sync()

	db, err := srv.DummyService().TestVoteData()
	if err != nil {
		t.Fatalf("failed to create base data: %v", err)
	}

	votingService := NewVoteService(db)

	user, err := srv.UserService().GetUser("email = ?", "t@t.t")
	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	event, err := srv.EventService().GetEvent(1)
	if err != nil {
		t.Fatalf("failed to fetch event: %v", err)
	}

	params := map[string]interface{}{
		"eventID": event.ID,
	}

	votes, err := votingService.GetVotes(params, 1)

	vote := votes[0]

	if err != nil {
		t.Fatalf("failed to fetch user: %v", err)
	}

	voteForm := forms.VoteInVotingForm{
		VoteID:       int(vote.ID),
		VoteOptionID: int(vote.Options[0].ID),
	}

	voteForm.VoteOptionID = int(vote.Options[1].ID)
	_, _ = votingService.Vote(voteForm, *user)

	voteForm.VoteOptionID = int(vote.Options[0].ID)
	_, err = votingService.Vote(voteForm, *user)

	assert.Error(t, err)

	vote.VoteType = models.CanChangeVote
	if err := db.Save(&vote).Error; err != nil {
		t.Fatalf("failed to save vote: %v", err)
	}

	_, err = votingService.Vote(voteForm, *user)
	assert.NoError(t, err)

	srv.Drop()
}

func TestClose(t *testing.T) {
	srv := New()

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
