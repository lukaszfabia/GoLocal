package email

import (
	"backend/internal/models"
	"backend/pkg"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/mail.v2"
)

// ContentProvider interface with methods to get email subject and info
type ContentProvider interface {
	GetSubject() string
	GetInfo() string
}

// Define types for email purposes
type ForgetPassword struct{}
type Verify struct{}

func (fp *ForgetPassword) GetSubject() string {
	return "Password Restore"
}

func (fp *ForgetPassword) GetInfo() string {
	return "It looks like you forgot your password. Don't worry, you're on the right path to restore your account."
}

func (v *Verify) GetSubject() string {
	return "Verify Your Account"
}

func (v *Verify) GetInfo() string {
	return "Almost there! Please use the code below to verify your account."
}

type Components struct {
	Username string
	Info     string
	Sub      string
	Code     string
}

func SendCode[T ContentProvider](email T, user models.User, code string) error {
	senderMail := os.Getenv("GMAIL_MAIL")
	senderPassword := os.Getenv("GMAIL_PASSWORD")

	if senderMail == "" || senderPassword == "" {
		return errors.New("SMTP credentials are missing")
	}

	username := fmt.Sprintf("%s %s", user.FirstName, user.LastName)

	m := mail.NewMessage()
	m.SetHeader("From", senderMail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Reply-To", "no-reply@example.com")
	m.SetAddressHeader("Cc", user.Email, username)
	m.SetHeader("Subject", email.GetSubject())

	body, err := pkg.ParseHTMLToString("base.html", Components{
		Username: username,
		Info:     email.GetInfo(),
		Sub:      email.GetSubject(),
		Code:     code,
	})

	if err != nil {
		log.Println("Error parsing email template:", err)
		return err
	}

	m.SetBody("text/html", body)

	d := mail.NewDialer("smtp.gmail.com", 587, senderMail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v", err)
		return errors.New("failed to send email")
	}

	return nil
}
