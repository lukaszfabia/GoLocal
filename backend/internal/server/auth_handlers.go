package server

import (
	"backend/internal/auth"
	"backend/internal/email"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/parsers"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"golang.org/x/crypto/bcrypt"
)

type providerKey string

const providerK providerKey = "provider"

func (s *Server) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.RefreshTokenRequest](r)

	if err != nil {
		s.NewResponse(w, http.StatusBadRequest, "Failed to decode form")
		return
	}

	// decode the refresh token
	sub, err := auth.DecodeJWT(form.Token)
	if err != nil {
		log.Println("Error decoding refresh token:", err)
		s.NewResponse(w, http.StatusUnauthorized, "Invalid or expired refresh token")
		return
	}

	// blacklist the old refresh token
	if err := s.db.TokenService().SetAsBlacklisted(form.Token); err != nil {
		log.Println("Error blacklisting token:", err)
		s.NewResponse(w, http.StatusInternalServerError, "Could not blacklist the token")
		return
	}

	// generate new tokens
	tokens, err := auth.GenerateJWT(sub, nil)
	if err != nil {
		log.Println("Error generating new tokens:", err)
		s.NewResponse(w, http.StatusInternalServerError, "Failed to generate new tokens")
		return
	}

	// send the new tokens
	s.NewResponse(w, http.StatusOK, *tokens)
}

func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.Register](r)

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	var user *models.User = &models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Password:  &form.Password,
		Email:     form.Email,
	}

	user, e := s.db.UserService().SaveUser(user)

	if e != nil {
		log.Println("Error saving user:", e)
		s.NewResponse(w, http.StatusBadRequest, "Could not create user account")
		return
	}

	if tokens, err := auth.GenerateJWT(user.ID, nil); err != nil {
		log.Println("Error generating JWT token:", err)
		s.NewResponse(w, http.StatusUnauthorized, err.Error())
	} else {

		s.NewResponse(w, http.StatusCreated, tokens)
	}

}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	provider, ok := r.Context().Value("provider").(string)

	// if logged in through a provider
	if ok && provider != "" {
		if err := gothic.Logout(w, r); err != nil {
			log.Println("Error during provider logout:", err)
			s.NewResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		s.NewResponse(w, http.StatusOK, "Successfully logged out!")
		return
	}

	// if not logged in through a provider, blacklist token
	authorization := r.Header.Get("Authorization")
	tokenStr := strings.TrimPrefix(authorization, "Bearer ")

	if err := s.db.TokenService().SetAsBlacklisted(tokenStr); err != nil {
		log.Println("Error blacklisting token:", err)
		s.NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// ctx := context.WithValue(r.Context(), userK, nil)
	// r = r.WithContext(ctx)

	s.NewResponse(w, http.StatusOK, "Successfully logged out!")
}

func (s *Server) Callback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	req := r.WithContext(context.WithValue(r.Context(), providerK, provider))

	user, err := gothic.CompleteUserAuth(w, req)
	if err != nil {
		s.NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var returnedUser = models.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		AuthProvider: &user.Provider,
		IsPremium:    false,
	}

	if avatar := user.AvatarURL; avatar != "" {
		returnedUser.AvatarURL = &avatar
	}

	if user, err := s.db.UserService().GetOrCreateUser(&returnedUser); err != nil {
		log.Println("Error creating user:", err)
		s.NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		tokens, err := auth.GenerateJWT(user.ID, nil)
		if err != nil {
			log.Println("Error generating JWT token:", err)
			s.NewResponse(w, http.StatusInternalServerError, "Error generating JWT token")
			return
		}

		s.NewResponse(w, http.StatusOK, *tokens)
	}

}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	if provider != "" {
		// callback returns tokens in the future
		req := r.WithContext(context.WithValue(context.Background(), providerK, provider))

		_, err := gothic.CompleteUserAuth(w, req)
		if err != nil {
			gothic.BeginAuthHandler(w, r)
			return
		}

	} else {
		form, err := parsers.DecodeJSON[forms.Login](r)

		if err != nil {
			s.InvalidFormResponse(w)
			return
		}

		user, err := s.db.UserService().GetUser(fmt.Sprintf("email = '%s'", form.Email))

		if err != nil {
			s.NewResponse(w, http.StatusNotFound, err.Error())
			return
		}

		comparePassword := func() error {
			if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(form.Password)); err != nil {
				return errors.New("credentials are invalid")
			}

			return nil
		}

		if tokens, err := auth.GenerateJWT(user.ID, &comparePassword); err != nil {
			s.NewResponse(w, http.StatusUnauthorized, err.Error())
		} else {
			s.NewResponse(w, http.StatusOK, *tokens)
		}
	}
}

func (s *Server) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// get user from ctx
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	code, err := s.store.SetCode(user.Email)

	if err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Something went wrong with storage")
		return
	}

	if err := email.SendCode(&email.Verify{}, *user, code); err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Failed to send code")
		return
	}

	s.NewResponse(w, http.StatusOK, "")
}

func (s *Server) VerifyCallbackHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.CodeRequest](r)
	user, ok := r.Context().Value("user").(*models.User)

	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	if !s.store.Compare(user.Email, form.Code) {
		s.NewResponse(w, http.StatusUnauthorized, "Failed to authorize, check spelling")
		return
	}

	user.IsVerified = true

	if _, err := s.db.UserService().SaveUser(user); err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Failed to save user")
		return
	}

	s.NewResponse(w, http.StatusOK, fmt.Sprintf("%s %s, you' ve been verifed", user.FirstName, user.LastName))
}

func (s *Server) PasswordResetCallbackHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.NewPasswordRequest](r)
	user, ok := r.Context().Value("user").(*models.User)

	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	user.Password = &form.Password

	_, err = s.db.UserService().SaveUser(user)

	if err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Could not change password")
		return
	}

	s.NewResponse(w, http.StatusOK, "Password changed!")
}

func (s *Server) PasswordResetHandler(w http.ResponseWriter, r *http.Request) {
	form, err := parsers.DecodeJSON[forms.RestoreAccount](r)

	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	user, err := s.db.UserService().GetUser(fmt.Sprintf("email = %s", form.Email))

	if err != nil {
		s.NewResponse(w, http.StatusNotFound, "Incorrect email")
		return
	}

	code, err := s.store.SetCode(user.Email)

	if err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Something went wrong with storage")
		return
	}

	if err := email.SendCode(&email.Verify{}, *user, code); err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Failed to send code")
		return
	}

	s.NewResponse(w, http.StatusOK, "")
}

func (s *Server) DeviceTokenRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: assign new device token to user
	form, err := parsers.DecodeJSON[forms.Device](r)

	if err != nil {
		s.NewResponse(w, http.StatusBadRequest, "Failed to decode form")
		return
	}

	if err := s.db.UserService().AddDevice(form); err != nil {
		s.NewResponse(w, http.StatusInternalServerError, "Failed to add device to user")
		return
	}

	s.NewResponse(w, http.StatusOK, "Successfully added new device!")
}
