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
		log.Println("No provided client")
		return errors.New("no provided client")
	}

	log.Println("Fetching device tokens for users:", n.usersIds)

	var devices []*models.DeviceToken

	if err := ns.db.Preload("Users", "users.id IN ?", n.usersIds).Find(&devices).Error; err != nil {
		log.Printf("Error fetching device tokens: %v", err)
		return err
	}

	log.Printf("Fetched %d device tokens", len(devices))

	// get only tokens
	tokens := []string{}

	for _, device := range devices {
		tokens = append(tokens, device.Token)
	}

	log.Printf("Collected %d tokens", len(tokens))

	notification := &messaging.Notification{
		Title: n.title,
		Body:  n.body,
	}

	if n.image != nil {
		notification.ImageURL = *n.image
		log.Printf("Notification image URL set: %s", *n.image)
	}

	message := &messaging.MulticastMessage{
		Tokens:       tokens,
		Notification: notification,
	}

	log.Println("Sending multicast message")

	batchResponse, err := ns.client.SendMulticast(context.Background(), message)
	if err != nil {
		log.Fatalf("Error sending multicast message: %v", err)
	}

	log.Printf("Successfully sent %d messages", batchResponse.SuccessCount)

	return nil
}
