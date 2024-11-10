package email

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

func sendEmail(w http.ResponseWriter, r *http.Request) {
	// Get the App Engine context from the request
	ctx := appengine.NewContext(r)

	// Create a new email message
	msg := &mail.Message{
		Sender:   "example@example.com",
		To:       []string{"recipient@example.com"},
		Subject:  "Hello from App Engine!",
		Body:     "This is a plain text email.",
		HTMLBody: "<html><body><p>This is an <b>HTML</b> email.</p></body></html>",
	}

	// Send the email
	if err := mail.Send(ctx, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Email sent successfully!"))
}
