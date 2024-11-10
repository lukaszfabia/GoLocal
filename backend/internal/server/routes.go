package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/gorilla/mux"
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
	api.HandleFunc("refresh-token/", s.RefreshTokenHandler).Methods(http.MethodPost)
	api.HandleFunc("/logout/", s.LogoutHandler).Methods(http.MethodGet)
	api.HandleFunc("/verify/", s.VerifyHandler).Methods(http.MethodPost)
	api.HandleFunc("/password-reset/", s.PasswordResetHandler).Methods(http.MethodPost)

	provider := api.PathPrefix("/{provider}").Subrouter()
	provider.HandleFunc("/login/", s.LoginHandler).Methods(http.MethodGet) // handles signing up
	provider.HandleFunc("/callback/", s.Callback).Methods(http.MethodGet)

	auth := api.PathPrefix("/auth").Subrouter()

	// setting middleware
	auth.Use(s.isAuth)
	// authenthicated endpoints
	auth.HandleFunc("/account/", s.AccountHandler).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	auth.HandleFunc("/logout/", s.LogoutHandler).Methods(http.MethodGet)

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
