package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "gorm.io/gorm"

	"github.com/coder/websocket"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Your API Title
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", s.HelloWorldHandler)

	r.HandleFunc("/ws", s.websocketHandler)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api").Subrouter()

	fileServer := http.FileServer(http.Dir("./media"))

	api.PathPrefix("/media/").Handler(http.StripPrefix("/api/media/", fileServer))

	// authentication routes
	api.HandleFunc("/login/", s.LoginHandler).Methods(http.MethodPost)
	api.HandleFunc("/sign-up/", s.SignUpHandler).Methods(http.MethodPost)
	api.HandleFunc("refresh-token/", s.RefreshTokenHandler).Methods(http.MethodPost)
	api.HandleFunc("/logout/", s.LogoutHandler).Methods(http.MethodGet)
	api.HandleFunc("/verify/", s.VerifyHandler).Methods(http.MethodPost)
	api.HandleFunc("/password-reset/", s.PasswordResetHandler).Methods(http.MethodPost)
	api.HandleFunc("/verify/callback/", s.VerifyCallbackHandler).Methods(http.MethodPost)
	api.HandleFunc("/password-reset/callback", s.PasswordResetCallbackHandler).Methods(http.MethodPost)
	api.HandleFunc("/device-token-registration/", s.DeviceTokenRegistrationHandler).Methods(http.MethodPost)

	provider := api.PathPrefix("/{provider}").Subrouter()
	provider.HandleFunc("/login/", s.LoginHandler).Methods(http.MethodGet) // handles signing up
	provider.HandleFunc("/callback/", s.Callback).Methods(http.MethodGet)

	auth := api.PathPrefix("/auth").Subrouter()
	// setting middleware
	// authenthicated endpoints
	auth.Use(s.isAuth)

	preference := api.PathPrefix("/preference").Subrouter()

	// preference survey routes
	preference.HandleFunc("/change-preference-survey", s.handleSurvey)
	preference.HandleFunc("/preference-survey/answer", s.handleSurveyAnswer)
	preference.HandleFunc("/preference-survey", s.getSurvey)

	// recommendation routes
	api.HandleFunc("/recommendation/{userID:[0-9]+}", s.getRecommendations)

	auth.HandleFunc("/account/", s.AccountHandler).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	auth.HandleFunc("/logout/", s.LogoutHandler).Methods(http.MethodGet)

	// events
	api.HandleFunc(`/{limit:[0-9]*}`, s.EventHandler).Methods(http.MethodGet)

	event := api.PathPrefix("/event").Subrouter()
	event.HandleFunc(`/{limit:[0-9]*}`, s.EventHandler).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	return r
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
