package auth

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 30 * 24 * 3600 // 30 days in seconds
	IsProd = true
)

func NewAuth() {
	// external providers
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// facebookClientID := os.Getenv("FACEBOOK_CLIENT_ID")
	// facebookClientSecret := os.Getenv("FACEBOOK_CLIENT_SECRET")

	port := os.Getenv("PORT")
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
			googleClientID,
			googleClientSecret,
			fmt.Sprintf("http://localhost:%s/api/google/callback/", port),
			"email",
			"profile",
		),

		// facebook.New(
		// 	facebookClientID,
		// 	facebookClientSecret,
		// 	fmt.Sprintf("http://localhost:%s/api/facebook/callback/", port),
		// ),
	)

}
