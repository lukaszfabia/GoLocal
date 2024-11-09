package pkg

import (
	"backend/internal/forms"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// register forms here
type Formable interface {
	forms.Login | forms.Register | forms.RefreshTokenRequest
}

func DecodeJSON[T Formable](r *http.Request) (*T, error) {
	form := new(T) // new instance of T
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, errors.New("Invalid JSON format")
	}

	return form, nil
}
