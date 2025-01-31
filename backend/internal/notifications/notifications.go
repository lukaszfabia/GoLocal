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

type Notification struct {
	Title    string
	Body     string
	Image    *string // optionally add image
	UsersIds []uint  // usrs group
	Author   uint
}

type NotificationService interface {
	SetClient(client MessagingClient)
	SendPush(n *Notification) error
}

type notificationServiceImpl struct {
	client MessagingClient
	db     *gorm.DB
}

func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationServiceImpl{
		client: nil,
		db:     db,
	}
}

func (ns *notificationServiceImpl) SetClient(client MessagingClient) {
	if client != nil {
		ns.client = client
		log.Println("Client has been set")
	} else {
		log.Println("Client is nil")
	}
}

func (ns *notificationServiceImpl) SendPush(n *Notification) error {
	if ns.client == nil {
		log.Println("No provided client")
		return errors.New("no provided client")
	}

	if len(n.UsersIds) < 1 {
		log.Println("No users specified")
		return errors.New("no users specified")
	}

	ids := functools.Filter(func(e uint) bool {
		return e != n.Author
	}, n.UsersIds)

	log.Println("Fetching device tokens for users:", ids)

	var devices []*models.DeviceToken

	if err := ns.db.
		Joins("JOIN user_devices ON user_devices.device_token_id = device_tokens.id").
		Joins("JOIN users ON users.id = user_devices.user_id").
		Where("users.id IN ? AND users.deleted_at IS NULL", ids).
		Find(&devices).Error; err != nil {
		log.Printf("Error fetching device tokens: %v", err)
		return err
	}

	log.Printf("Fetched %d device tokens", len(devices))

	tokens := make([]string, 0, len(devices))
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
		log.Printf("Error sending multicast message: %v", err)
		return err
	}

	log.Printf("Successfully sent %d messages", batchResponse.SuccessCount)

	return nil
}
