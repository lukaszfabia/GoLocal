package auth_test

import (
	"backend/internal/auth"
	"testing"
)

func JWTest(t *testing.T) {
	var userID uint = 1

	tokens, err := auth.GenerateJWT(userID, nil)
	if err != nil {
		t.Errorf("Failed to generate JWT: %v", err)
	}

	if _, err := auth.DecodeJWT(tokens.Access); err != nil {
		t.Errorf("Failed to decode Access token: %v", err)
	}

	sub, err := auth.DecodeJWT(tokens.Refresh)
	if err != nil {
		t.Errorf("Failed to decode Refresh token: %v", err)
	}
	if sub != userID {
		t.Errorf("Expected subject %v, got %v", userID, sub)
	}
}
