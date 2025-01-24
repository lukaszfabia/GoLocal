package database

import (
	"backend/internal/models"
	"backend/internal/recommendation"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var allModels []any = []any{
	&models.User{},
	&models.Location{},
	&models.Coords{},
	&models.Address{},
	&models.Event{},
	&models.Comment{},

	&models.Vote{},
	&models.VoteOption{},
	&models.VoteAnswer{},

	&models.Opinion{},
	&models.Tag{},

	&models.PreferenceSurvey{},
	&models.PreferenceSurveyQuestion{},
	&models.PreferenceSurveyAnswer{},
	&models.PreferenceSurveyOption{},
	&models.PreferenceSurveyAnswerOption{},

	&models.Recommendation{},
	&models.BlacklistedTokens{},
	&models.DeviceToken{},
}

var (
	database   = strings.TrimSpace(os.Getenv("DB_DATABASE"))
	password   = strings.TrimSpace(os.Getenv("DB_PASSWORD"))
	username   = strings.TrimSpace(os.Getenv("DB_USERNAME"))
	port       = strings.TrimSpace(os.Getenv("DB_PORT"))
	host       = strings.TrimSpace(os.Getenv("DB_HOST"))
	dbInstance *service
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Make table migration to database.
	// If something go wrong, info is logeed in console.
	Sync()

	// Fill database with dummy data
	DummyService() DummyService

	UserService() UserService
	PreferenceSurveyService() PreferenceSurveyService
	RecommendationService() recommendation.RecommendationService
	EventService() EventService
	TokenService() TokenService
	VoteService() VoteService
	NotificationService() NotificationService
}

type service struct {
	db *gorm.DB

	userService             UserService
	preferenceSurveyService PreferenceSurveyService
	recommendationService   recommendation.RecommendationService
	tokenService            TokenService
	dummyService            DummyService
	notificationService     NotificationService

	eventService EventService
	voteService  VoteService
}

// Initializes new database connection and services
// Returns service
func New() Service {
	// Reuse Connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		log.Printf("Provided dsn %s\n", dsn)
		panic("Can't connect to db!\n" + err.Error())
	} else {
		log.Println("Successfully connected to db!")

		// register services

		userService := NewUserService(db)
		tokenService := NewTokenService(db)
		dummyService := NewDummyService(db)
		prefenceSurveyService := NewPreferenceSurveyService(db)
		recommendationService := recommendation.NewRecommendationService(db)
		eventService := NewEventService(db)
		voteService := NewVoteService(db)
		notificationService := NewNotificationService(db)

		return &service{
			db:                      db,
			userService:             userService,
			tokenService:            tokenService,
			dummyService:            dummyService,
			preferenceSurveyService: prefenceSurveyService,
			recommendationService:   recommendationService,
			eventService:            eventService,
			voteService:             voteService,
			notificationService:     notificationService,
		}
	}
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	var err error
	var sqlDB *sql.DB

	sqlDB, err = s.db.DB()

	if err != nil {
		log.Println(err)
	}

	// Ping the database
	err = sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatal(stats["error"]) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)

	if dbInstance, err := s.db.DB(); err != nil {
		log.Println(err)
		return err
	} else {
		_ = dbInstance.Close()
	}

	return nil
}

// Migrates tables using allModels to database
func (s *service) Sync() {
	for _, model := range allModels {
		if err := s.db.AutoMigrate(model); err != nil {
			log.Println(err)
		}
	}
	log.Println("Migrating models has been done.")

	s.dummyService.Cook()
}

func (s *service) RecommendationService() recommendation.RecommendationService {
	return s.recommendationService
}
