package email_test

import (
	"backend/internal/email"
	"backend/internal/models"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendCode(t *testing.T) {

	godotenv.Load("../../.env")

	err := email.SendCode(&email.ForgetPassword{}, models.User{
		FirstName: "Lukasz",
		LastName:  "Fabia",
		Email:     "ufabia03@gmail.com",
	}, "82734")

	if err != nil {
		t.Error(err)
	}
}
