package database

import (
	"backend/internal/models"
	"context"
	"errors"
	"log"

	"firebase.google.com/go/messaging"
	"gorm.io/gorm"
)

// Notification model which is sent to groups of users
type notification struct {
	title    string
	body     string
	image    *string // optionally add image
	usersIds []uint  // usrs group
}

func NewNotification(title, body string, image *string, usersIds []uint) notification {
	return notification{
		title:    title,
		body:     body,
		image:    image,
		usersIds: usersIds,
	}
}

type NotificationService interface {
	// Set client before you call send push notification
	SetClient(client *messaging.Client)
	SendPush(n *notification) error
}

type notificationServiceImpl struct {
	client *messaging.Client
	db     *gorm.DB
}

func (s *service) NotificationService() NotificationService {
	return s.notificationService
}

func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationServiceImpl{
		client: nil,
		db:     db,
	}
}

func (ns *notificationServiceImpl) SetClient(client *messaging.Client) {
	ns.client = client
	log.Println("Client has been set")
}

func (ns *notificationServiceImpl) SendPush(n *notification) error {
	if ns.client == nil {
		return errors.New("no provided client")
	}

	var devices []*models.DeviceToken

	if err := ns.db.Preload("Users", "users.id IN ?", n.usersIds).Find(&devices).Error; err != nil {
		return err
	}
	// get only tokens
	tokens := []string{}

	for _, device := range devices {
		tokens = append(tokens, device.Token)
	}

	notification := &messaging.Notification{
		Title: n.title,
		Body:  n.body,
	}

	if n.image != nil {
		notification.ImageURL = *n.image
	}

	message := &messaging.MulticastMessage{
		Tokens:       tokens,
		Notification: notification,
	}

	batchResponse, err := ns.client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending multicast message: %v", err)
	}

	log.Printf("Successfully sent %d messages", batchResponse.SuccessCount)

	return nil
}
