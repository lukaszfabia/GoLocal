package server

import (
	"backend/internal/server/account"
	"backend/internal/server/event"
	recommendation_handler "backend/internal/server/recommendation"
	"backend/internal/server/vote"
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

	// newwww
	eventHandler := event.EventHandler{
		EventService:        s.db.EventService(),
		NotificationService: s.db.NotificationService(),
	}

	voteHandler := vote.VoteHandler{
		VoteService: s.db.VoteService(),
		UserService: s.db.UserService(),
	}

	accountHandler := account.AccountHandler{
		UserService: s.db.UserService(),
	}

	recommendationHandler := recommendation_handler.RecommendationHandler{
		UserService:           s.db.UserService(),
		EventService:          s.db.EventService(),
		RecommendationService: s.db.RecommendationService(),
	}

	// newwww

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

	preference := auth.PathPrefix("/preference").Subrouter()

	// preference survey routes
	preference.HandleFunc("/change-preference-survey", s.handleSurvey)
	preference.HandleFunc("/preference-survey/answer", s.handleSurveyAnswer)
	preference.HandleFunc("/preference-survey", s.getSurvey)
	preference.HandleFunc("/was-survey-filled", s.hasAccessToRecommendedEvents)

	// recommendation routes
	auth.HandleFunc("/recommendations", recommendationHandler.Handle).
		Methods(http.MethodGet)

	auth.HandleFunc("/account/", accountHandler.Handle).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	auth.HandleFunc("/logout/", s.LogoutHandler).Methods(http.MethodGet)

	// events to remove
	api.HandleFunc(`/{limit:[0-9]*}`, s.EventHandler).Methods(http.MethodGet)

	// votes
	api.HandleFunc(`/{limit:[0-9]*}`, s.VoteHandler).Methods(http.MethodGet)

	//-------------------------------------Event-------------------------------------//

	event := auth.PathPrefix("/event").Subrouter()

	promoRouter := event.PathPrefix("/{id}/promo").Subrouter()
	promoRouter.Use(eventHandler.ValidatePromoteEvent)
	promoRouter.HandleFunc("", eventHandler.PromoteEvent)

	reportRouter := event.PathPrefix("/report").Subrouter()
	reportRouter.Use(eventHandler.ValidateReportEvent)
	reportRouter.HandleFunc("", eventHandler.ReportEvent)

	defaultRouter := event.PathPrefix("").Subrouter()
	defaultRouter.Use(eventHandler.Validate)
	defaultRouter.HandleFunc(`/{limit:[0-9]*}`, eventHandler.Handle).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	//-------------------------------------Vote-------------------------------------//

	vote := auth.PathPrefix("/vote").Subrouter()
	vote.HandleFunc(`/{limit:[0-9]*}`, voteHandler.Handle).
		Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	health := s.db.Health()
	log.Println(s.db.Health())

	if health["status"] == "down" {
		s.NewResponse(w, http.StatusInternalServerError, nil)
		return
	}

	s.NewResponse(w, http.StatusOK, health["message"])
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
