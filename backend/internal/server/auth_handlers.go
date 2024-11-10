package server

import (
	"backend/internal/auth"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg"
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

func (s *Server) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	form, err := pkg.DecodeJSON[forms.RefreshTokenRequest](r)

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
	form, err := pkg.DecodeJSON[forms.Register](r)

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
		s.NewResponse(w, http.StatusBadRequest, "Could not create user account")
		return
	}

	if tokens, err := auth.GenerateJWT(user.ID, nil); err != nil {
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

	ctx := context.WithValue(r.Context(), "user", nil)
	r = r.WithContext(ctx)

	s.NewResponse(w, http.StatusOK, "Successfully logged out!")
}

func (s *Server) Callback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	req := r.WithContext(context.WithValue(r.Context(), "provider", provider))

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
		req := r.WithContext(context.WithValue(context.Background(), "provider", provider))

		_, err := gothic.CompleteUserAuth(w, req)
		if err != nil {
			gothic.BeginAuthHandler(w, r)
			return
		}

	} else {
		form, err := pkg.DecodeJSON[forms.Login](r)

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

// Send code on e-mail
func (s *Server) sendCode(w http.ResponseWriter, r *http.Request) {
	type emailBody struct {
		Email string `json:"email"`
	}

	_, err := pkg.DecodeJSON[forms.VerifyUser](r)

	if err != nil {
		s.InvalidFormResponse(w)
	}

	// send email
}

func (s *Server) VerifyHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) PasswordReset(w http.ResponseWriter, r *http.Request) {

}
