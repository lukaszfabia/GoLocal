package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"

	"backend/internal/database"
	"backend/internal/store"
)

type Server struct {
	port int

	db    database.Service
	store store.Store
}

type response struct {
	HTTPCode int `json:"status"`
	Data     any `json:"data,omitempty"`
}

// Creates new server response
func (s *Server) NewResponse(w http.ResponseWriter, httpCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpCode)

	err := json.NewEncoder(w).Encode(&response{
		HTTPCode: httpCode,
		Data:     data,
	})

	if err != nil {
		log.Printf("Error encoding response: %v", err)
		if httpCode == http.StatusOK {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func (s *Server) InvalidFormResponse(w http.ResponseWriter) {
	s.NewResponse(w, http.StatusBadRequest, "Invalid body")
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(strings.TrimSpace(os.Getenv("PORT")))

	opt := option.WithCredentialsFile("./golocal-firebase.json")

	fApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Failed to init firebase app: %s", err.Error())
	}

	client, err := fApp.Messaging(context.Background())

	if err != nil {
		log.Println(err.Error())
	}

	NewServer := &Server{
		port: port,

		db:    database.New(),
		store: store.New(),
	}

	NewServer.db.NotificationService().SetClient(client)

	NewServer.db.Sync()

	// Declare Server config
	server := &http.Server{
		Addr:         ":8080",
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
