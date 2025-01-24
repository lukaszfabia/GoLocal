package store

import (
	"testing"
)

func TestSetCode(t *testing.T) {
	store := New()

	// Test: SetCode should generate a code and store it
	email := "test@example.com"
	code, err := store.SetCode(email)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(code) != 6 {
		t.Errorf("expected code length of 6, got %d", len(code))
	}

	// Test: Code should be stored correctly
	if !store.Compare(email, code) {
		t.Errorf("stored code doesn't match the expected code")
	}
}

func TestCompareInvalidEmail(t *testing.T) {
	store := New()

	// Test: Compare with an invalid email
	code := "123456"
	email := "test@example.com"
	if store.Compare(email, code) {
		t.Errorf("expected false for invalid email comparison")
	}
}

func TestCompareInvalidCode(t *testing.T) {
	store := New()

	// Test: SetCode and Compare with wrong code
	email := "test@example.com"
	expectedCode, _ := store.SetCode(email)

	// Compare with incorrect code
	if store.Compare(email, "wrongcode") {
		t.Errorf("expected false for incorrect code comparison")
	}

	// Compare with the correct code
	if !store.Compare(email, expectedCode) {
		t.Errorf("expected true for correct code comparison")
	}
}

func TestClear(t *testing.T) {
	store := New()

	// Test: Set a code and then clear it
	email := "test@example.com"
	code, _ := store.SetCode(email)

	// Ensure the code is stored
	if !store.Compare(email, code) {
		t.Errorf("expected stored code to match")
	}

	// Clear the code
	store.clear(email)

	// Ensure the code is cleared
	if store.Compare(email, code) {
		t.Errorf("expected code to be cleared")
	}
}

func TestGenerateCodeLength(t *testing.T) {
	// Test: GenerateCode should always return a 6 digit code
	code, err := generateCode(6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(code) != 6 {
		t.Errorf("expected code length of 6, got %d", len(code))
	}

	// Test: GenerateCode should fail with length < 6
	_, err = generateCode(5)
	if err == nil {
		t.Errorf("expected error for length < 6")
	}
}
