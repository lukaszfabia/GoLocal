package parsers_test

import (
	"backend/internal/forms"
	"backend/pkg/parsers"
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testStruct struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Job      bool   `json:"job"`
	Age      int    `json:"age"`
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name         string
		inputDate    string
		expectedDate time.Time
		expectError  bool
	}{
		{
			name:         "Valid date",
			inputDate:    "2025-01-24",
			expectedDate: time.Date(2025, time.January, 24, 0, 0, 0, 0, time.UTC),
			expectError:  false,
		},
		{
			name:         "Invalid date format",
			inputDate:    "24-01-2025", // Invalid format
			expectedDate: time.Time{},  // Zero value
			expectError:  true,
		},
		{
			name:         "Empty date string",
			inputDate:    "",
			expectedDate: time.Time{}, // Zero value
			expectError:  true,
		},
		{
			name:         "Leap year date",
			inputDate:    "2024-02-29",
			expectedDate: time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			expectError:  false,
		},
		{
			name:         "Non-leap year February 29th",
			inputDate:    "2023-02-29",
			expectedDate: time.Time{}, // Zero value since it's an invalid date
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parsers.ParseDate(tt.inputDate)

			// Check if result is expected date
			if !result.Equal(tt.expectedDate) {
				t.Errorf("expected %v, got %v", tt.expectedDate, result)
			}

			// Check for error cases (zero time value indicates an error)
			if (result.IsZero() && !tt.expectError) || (!result.IsZero() && tt.expectError) {
				t.Errorf("expected error: %v, got: %v", tt.expectError, result.IsZero())
			}
		})
	}
}

func TestDecodeMultipartForm(t *testing.T) {
	tests := []struct {
		name          string
		setupRequest  func() *http.Request
		expectedForm  *testStruct
		expectedError error
	}{
		{
			name: "Valid form data",
			setupRequest: func() *http.Request {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				writer.WriteField("name", "John")
				writer.WriteField("lastname", "Doe")
				writer.WriteField("job", "true")
				writer.WriteField("age", "30")
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/test", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				return req
			},
			expectedForm: &testStruct{
				Name:     "John",
				Lastname: "Doe",
				Job:      true,
				Age:      30,
			},
			expectedError: nil,
		},
		{
			name: "Invalid form data (invalid field)",
			setupRequest: func() *http.Request {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				writer.WriteField("name", "John")
				writer.WriteField("lastname", "Doe")
				writer.WriteField("job", "true")
				writer.WriteField("invalid_field", "data")
				writer.Close()

				req := httptest.NewRequest(http.MethodPost, "/test", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				return req
			},
			expectedForm:  nil,
			expectedError: errors.New("invalid form data"),
		},
		{
			name: "Error parsing form (empty form)",
			setupRequest: func() *http.Request {
				body := &bytes.Buffer{}
				req := httptest.NewRequest(http.MethodPost, "/test", body)
				req.Header.Set("Content-Type", "multipart/form-data")
				return req
			},
			expectedForm:  nil,
			expectedError: errors.New("invalid form format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setupRequest()
			form, err := parsers.DecodeMultipartForm[testStruct](req)

			// Check for errors
			if (err != nil && tt.expectedError == nil) || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}

			// Check form data
			if form != nil && *form != *tt.expectedForm {
				t.Errorf("expected form: %v, got: %v", tt.expectedForm, form)
			}
		})
	}
}

// helper function to create a request with JSON body
func createRequestWithBody(jsonData []byte) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestDecodeJSON(t *testing.T) {
	tests := []struct {
		name          string
		setupRequest  func() *http.Request
		expectedForm  *forms.Login
		expectedError string // Change to string for easier error comparison
	}{
		{
			name: "Valid JSON",
			setupRequest: func() *http.Request {
				data := `{"email":"joe.doe@example.com", "password":"P@ssw0rd"}`
				return createRequestWithBody([]byte(data))
			},
			expectedForm: &forms.Login{
				Email:    "joe.doe@example.com",
				Password: "P@ssw0rd",
			},
			expectedError: "",
		},
		{
			name: "Invalid JSON format",
			setupRequest: func() *http.Request {
				// Invalid JSON (missing closing brace)
				data := `{"email":"joe.doe@example.com", "password":"P@ssw0rd"`
				return createRequestWithBody([]byte(data))
			},
			expectedForm:  nil,
			expectedError: "invalid JSON format",
		},
		{
			name: "Empty request body",
			setupRequest: func() *http.Request {
				// Empty JSON
				data := `{}`
				return createRequestWithBody([]byte(data))
			},
			expectedForm:  &forms.Login{}, // Empty struct should be returned
			expectedError: "",
		},
		{
			name: "No JSON content",
			setupRequest: func() *http.Request {
				// Non-JSON content (plain text)
				data := `Hello, this is not JSON`
				return createRequestWithBody([]byte(data))
			},
			expectedForm:  nil,
			expectedError: "invalid JSON format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setupRequest()
			form, err := parsers.DecodeJSON[forms.Login](req)

			// Check for expected error
			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			} else if err == nil && tt.expectedError != "" {
				t.Errorf("expected error: %v, but got none", tt.expectedError)
			}

			// Check form fields for equality
			if form != nil && *form != *tt.expectedForm {
				t.Errorf("expected form: %v, got: %v", tt.expectedForm, form)
			}
		})
	}
}

func TestParseURLQuery(t *testing.T) {
	tests := []struct {
		name           string
		setupRequest   func() *http.Request
		expectedResult map[string]any
		expectedError  error
	}{
		{
			name: "Valid query parameters",
			setupRequest: func() *http.Request {
				data := "/?name=John&lastname=Doe&job=true&age=30"
				req, _ := http.NewRequest(http.MethodGet, data, nil)
				return req
			},
			expectedResult: map[string]any{
				"name":     "John",
				"lastname": "Doe",
				"job":      true,
				"age":      30,
			},
			expectedError: nil,
		},
		{
			name: "Missing some query parameters",
			setupRequest: func() *http.Request {
				data := "/?name=John&age=30"
				req, _ := http.NewRequest(http.MethodGet, data, nil)
				return req
			},
			expectedResult: map[string]any{
				"name": "John",
				"age":  30,
			},
			expectedError: nil,
		},
		{
			name: "Empty query parameters",
			setupRequest: func() *http.Request {
				data := "/?"
				req, _ := http.NewRequest(http.MethodGet, data, nil)
				return req
			},
			expectedResult: map[string]any{},
			expectedError:  nil,
		},
		{
			name: "Custom query parameters",
			setupRequest: func() *http.Request {
				data := "/?name=John&age=30&extraParam=extraValue"
				req, _ := http.NewRequest(http.MethodGet, data, nil)
				return req
			},
			expectedResult: map[string]any{
				"name": "John",
				"age":  30,
			},
			expectedError: nil,
		},
		{
			name: "Nonexistent query parameters",
			setupRequest: func() *http.Request {
				data := "/?nonexistent=1"
				req, _ := http.NewRequest(http.MethodGet, data, nil)
				return req
			},
			expectedResult: map[string]any{},
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setupRequest()
			result := parsers.ParseURLQuery(req, testStruct{})

			// Check errors (since ParseURLQuery doesn't return an error, just an empty map)
			if len(result) != len(tt.expectedResult) {
				t.Errorf("expected result: %v, got: %v", tt.expectedResult, result)
			}

			// Check individual parameters
			for key, expectedValue := range tt.expectedResult {
				if result[key] != expectedValue {
					t.Errorf("expected %v for key %s, got %v", expectedValue, key, result[key])
				}
			}
		})
	}
}
