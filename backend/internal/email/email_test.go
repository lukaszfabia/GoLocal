package email_test

import (
	"backend/internal/email"
	"backend/internal/models"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendCode(t *testing.T) {

	godotenv.Load("../../.env")

	gmail := os.Getenv("GMAIL_MAIL")
	if gmail == "" {
		gmail = "ufabia03@gmail.com"
	}

	err := email.SendCode(&email.ForgetPassword{}, models.User{
		FirstName: "Lukasz",
		LastName:  "Fabia",
		Email:     gmail,
	}, "82734")

	if err != nil {
		t.Error(err)
	}
}
