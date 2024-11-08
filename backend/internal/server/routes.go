package server

import (
	"backend/internal/auth"
	"backend/internal/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"golang.org/x/crypto/bcrypt"

	"github.com/coder/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", s.HelloWorldHandler)

	r.HandleFunc("/ws", s.websocketHandler)

	api := r.PathPrefix("/api").Subrouter()
	// authentication routes
	api.HandleFunc("/login/", s.LoginHandler).Methods(http.MethodPost)
	api.HandleFunc("/sign-up/", s.SignUpHandler).Methods(http.MethodPost)
	api.HandleFunc("/logout/", s.Logout)

	provider := api.PathPrefix("/{provider}").Subrouter()
	provider.HandleFunc("/login/", s.LoginHandler).Methods(http.MethodGet) // handles signing up
	provider.HandleFunc("/callback/", s.Callback).Methods(http.MethodGet)

	auth := api.PathPrefix("/auth").Subrouter()

	// setting middleware
	auth.Use(s.isAuth)
	// authenthicated endpoints
	auth.HandleFunc("/account/", s.AccountHandler).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	auth.HandleFunc("/logout/", s.Logout).Methods(http.MethodGet)

	return r
}

func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	provider := r.Context().Value("provider")

	// if he was logged by provider
	if provider != "" {
		if err := gothic.Logout(w, r); err != nil {
			log.Println(err)

			s.NewResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		s.NewResponse(w, http.StatusOK, "Successfully logged out!")
	} else {
		authorization := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authorization, "Bearer ")

		if err := s.db.TokenService().SetAsBlacklisted(tokenStr); err != nil {
			log.Println("Can not blacklist token")
			s.NewResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		removeUser := context.WithValue(r.Context(), "user", "nil")
		r.WithContext(removeUser)

		s.NewResponse(w, http.StatusOK, map[string]any{
			"message": "Logged out!",
		})
	}

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
		AuthProvider: user.Provider,
		IsPremium:    false,
	}

	if avatar := user.AvatarURL; avatar != "" {
		returnedUser.AvatarURL = &avatar
	}

	if user, err := s.db.UserService().GetOrCreateUser(&returnedUser); err != nil {
		log.Println("Error creating user:", err)
		s.NewResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		token, err := auth.GenerateJWT(user.ID, nil)
		if err != nil {
			log.Println("Error generating JWT token:", err)
			s.NewResponse(w, http.StatusInternalServerError, "Error generating JWT token")
			return
		}

		s.NewResponse(w, http.StatusOK, map[string]interface{}{
			"message": "User authenticated successfully",
			"token":   token,
		})
	}

}

func (s *Server) AccountHandler(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]

	if provider != "" {
		req := r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothUser, err := gothic.CompleteUserAuth(w, req)
		if err != nil {
			gothic.BeginAuthHandler(w, r)
			return
		}

		s.NewResponse(w, http.StatusOK, map[string]any{
			"message": "Successfully logged in",
			"user":    gothUser,
		})
	} else {
		var lf models.LoginForm
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&lf)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			s.NewResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := s.db.UserService().GetUser(fmt.Sprintf("email = %s", lf.Email))

		if err != nil {
			s.NewResponse(w, http.StatusNotFound, err.Error())
			return
		}

		comparePassword := func() error {
			if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(lf.Password)); err != nil {
				return errors.New("credentials are invalid")
			}

			return nil
		}

		if token, err := auth.GenerateJWT(user.ID, &comparePassword); err != nil {
			s.NewResponse(w, http.StatusOK, map[string]any{
				"message": "Successfully logged in!",
				"token":   token,
			})
		}
	}
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	type testRequest struct {
		Message string `json:"message"`
		Year    int    `json:"year"`
	}

	req := testRequest{
		Message: "test",
		Year:    time.Now().Year(),
	}

	s.NewResponse(w, http.StatusOK, req)
}

func (s *Server) healthHandler(w http.ResponseWriter, _ *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
