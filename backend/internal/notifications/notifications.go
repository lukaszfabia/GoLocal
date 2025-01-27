package notifications

import (
	"backend/internal/models"
	"backend/pkg/functools"
	"context"
	"errors"
	"log"

	"firebase.google.com/go/messaging"
	"gorm.io/gorm"
)

// Notification model which is sent to groups of users
type Notification struct {
	Title    string
	Body     string
	Image    *string // optionally add image
	UsersIds []uint  // usrs group
	Author   uint
}

type NotificationService interface {
	// Set client before you call send push notification
	SetClient(client *messaging.Client)
	SendPush(n *Notification) error
}

type notificationServiceImpl struct {
	client *messaging.Client
	db     *gorm.DB
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

func (ns *notificationServiceImpl) SendPush(n *Notification) error {
	if ns.client == nil {
		log.Println("No provided client")
		return errors.New("no provided client")
	}

	// remove here author
	ids := functools.Filter(func(e uint) bool {
		return e != n.Author
	}, n.UsersIds)

	log.Println("Fetching device tokens for users:", ids)

	var devices []*models.DeviceToken

	if err := ns.db.Preload("Users", "users.id IN ?", ids).Find(&devices).Error; err != nil {
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
		Title: n.Title,
		Body:  n.Body,
	}

	if n.Image != nil {
		notification.ImageURL = *n.Image
		log.Printf("Notification image URL set: %s", *n.Image)
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
