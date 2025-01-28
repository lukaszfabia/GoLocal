package email_test

// func TestSendCode(t *testing.T) {

// 	godotenv.Load("../../.env")

// 	gmail := os.Getenv("GMAIL_MAIL")
// 	if gmail == "" {
// 		gmail = "ufabia03@gmail.com"
// 	}

// 	err := email.SendCode(&email.ForgetPassword{}, models.User{
// 		FirstName: "Lukasz",
// 		LastName:  "Fabia",
// 		Email:     gmail,
// 	}, "82734")

// 	if err != nil {
// 		t.Error(err)
// 	}
// }
