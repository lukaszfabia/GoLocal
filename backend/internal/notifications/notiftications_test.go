package notifications_test

import (
	"backend/internal/notifications"
	"context"
	"fmt"
	"testing"

	"firebase.google.com/go/messaging"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mockMessagingClient struct {
	sendMulticastFunc func(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
}

func (m *mockMessagingClient) SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error) {
	return m.sendMulticastFunc(ctx, message)
}

func TestSendPush_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.26"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM connection: %v", err)
	}

	img := "http://tomaszAndrzejDziałowy.jpg"

	mockClient := &mockMessagingClient{
		sendMulticastFunc: func(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error) {
			assert.Equal(t, "Test Title", message.Notification.Title)
			assert.Equal(t, "Test Body", message.Notification.Body)
			assert.Equal(t, img, message.Notification.ImageURL)
			assert.ElementsMatch(t, []string{"token1", "token2"}, message.Tokens)

			return &messaging.BatchResponse{SuccessCount: len(message.Tokens)}, nil
		},
	}

	service := notifications.NewNotificationService(gormDB)
	service.SetClient(mockClient)

	deviceRows := sqlmock.NewRows([]string{"id", "token", "os_version", "platform"}).
		AddRow(1, "token1", "14.5", "iOS").
		AddRow(2, "token2", "11.0", "Android")

	expectedQuery := `SELECT .* FROM .* WHERE .*`
	mock.ExpectQuery(expectedQuery).WithArgs(2).WillReturnRows(deviceRows)

	notification := &notifications.Notification{
		Title:    "Test Title",
		Body:     "Test Body",
		Image:    &img,
		UsersIds: []uint{1, 2},
		Author:   1,
	}

	err = service.SendPush(notification)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendPush_DatabaseQueryFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.26"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM connection: %v", err)
	}

	mockClient := &mockMessagingClient{}

	service := notifications.NewNotificationService(gormDB)
	service.SetClient(mockClient)

	expectedQuery := `SELECT .* FROM .* WHERE .*`
	mock.ExpectQuery(expectedQuery).WithArgs(2).WillReturnError(fmt.Errorf("database error"))

	notification := &notifications.Notification{
		Title:    "Test Title",
		Body:     "Test Body",
		UsersIds: []uint{1, 2},
		Author:   1,
	}

	err = service.SendPush(notification)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendPush_FCMClientFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.26"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM connection: %v", err)
	}

	mockClient := &mockMessagingClient{
		sendMulticastFunc: func(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error) {
			return nil, fmt.Errorf("FCM client error")
		},
	}

	service := notifications.NewNotificationService(gormDB)
	service.SetClient(mockClient)

	deviceRows := sqlmock.NewRows([]string{"id", "token", "os_version", "platform"}).
		AddRow(1, "token1", "14.5", "iOS").
		AddRow(2, "token2", "11.0", "Android")

	expectedQuery := `SELECT .* FROM .* WHERE .*`
	mock.ExpectQuery(expectedQuery).WithArgs(2).WillReturnRows(deviceRows)

	notification := &notifications.Notification{
		Title:    "Test Title",
		Body:     "Test Body",
		UsersIds: []uint{1, 2},
		Author:   1,
	}

	err = service.SendPush(notification)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FCM client error")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendPush_InvalidInput(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.26"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM connection: %v", err)
	}

	mockClient := &mockMessagingClient{}

	service := notifications.NewNotificationService(gormDB)
	service.SetClient(mockClient)

	notification := &notifications.Notification{
		Title:    "Test Title",
		Body:     "Test Body",
		UsersIds: []uint{},
		Author:   1,
	}

	err = service.SendPush(notification)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no users specified")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendPush_NoClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.26"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening GORM connection: %v", err)
	}

	service := notifications.NewNotificationService(gormDB)

	notification := &notifications.Notification{
		Title:    "Test Title",
		Body:     "Test Body",
		UsersIds: []uint{1, 2},
		Author:   1,
	}

	err = service.SendPush(notification)

	assert.Error(t, err)

	assert.Contains(t, err.Error(), "no provided client")

	assert.NoError(t, mock.ExpectationsWereMet())
}
