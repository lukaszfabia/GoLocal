package notifications

import (
	"context"

	"firebase.google.com/go/messaging"
)

// sick
type MessagingClient interface {
	SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
}
