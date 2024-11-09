package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type User struct {
	Model
	FirstName    string  `gorm:"not null:size:255" json:"firstName"`
	LastName     string  `gorm:"not null:size:255" json:"lastName"`
	Email        string  `gorm:"not null;size:100;unique" json:"email"`
	Password     *string `gorm:"null;size:400" json:"-"`
	AuthProvider *string `gorm:"null" json:"provider"`

	IsPremium bool    `gorm:"default:false" json:"isPremium"`
	AvatarURL *string `gorm:"null;size:1024" json:"avatarUrl"`

	Followers []*User `gorm:"many2many:user_followers" json:"followers"`
	Following []*User `gorm:"many2many:user_following" json:"following"`

	Comments []*Comment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"comments"`
	Votes    []*Vote    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"votes"`
}

type Location struct {
	Model
	City    string `gorm:"size:100" json:"city"`
	Country string `gorm:"size:100" json:"country"`
	Zip     string `gorm:"size:20" json:"zip"`

	Coords   *Coords `gorm:"foreignKey:CoordsID;references:ID;constraint:OnDelete:CASCADE" json:"coords"`
	CoordsID uint    `json:"coordsID"`

	Address   *Address `gorm:"foreignKey:AddressID;references:ID;constraint:OnDelete:CASCADE" json:"address"`
	AddressID uint     `json:"addressID"`
}

type Coords struct {
	Model
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
}

type Address struct {
	Model
	Street         string `gorm:"not null;size:255" json:"street"`
	StreetNumber   int    `gorm:"null" json:"streetNumber"`
	AdditionalInfo string `gorm:"null:size:512" json:"additionalInfo"`
}

type Event struct {
	Model
	EventOrganizers []*User `gorm:"many2many:event_organizers" json:"eventOrganizers"`

	Title       string  `gorm:"not null;size:255" json:"title"`
	Description string  `gorm:"default:'';null;size:255" json:"description"`
	IamgeURL    *string `gorm:"null;size:1024" json:"image"`
	IsAdultOnly bool    `gorm:"default:true" json:"isAdultOnly"`

	// Timestamp with time zone
	StartDate  *time.Time `gorm:"type:date;not null" json:"startDate"`
	FinishDate *time.Time `gorm:"type:date;" json:"finishDate"`

	Location   *Location `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE" json:"location"`
	LocationID uint      `json:"locationID"`

	Comments []*Comment `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"comments"`
	Votes    []*Vote    `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"votes"`
}

type Comment struct {
	Model
	UserID  uint   `json:"userID"`
	EventID uint   `json:"eventID"`
	Content string `gorm:"not null;size:255" json:"content"`
}

type Vote struct {
	Model
	UserID  uint                `json:"userID"`
	EventID uint                `json:"eventID"`
	State   ParticipationStatus `gorm:"type:text;not null" json:"state"`
}

type Survey struct {
	Model
	Event   *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:RESTRICT" json:"event"`
	EventID uint   `json:"eventID"`

	Questions []*SurveyQuestion `gorm:"foreignKey:SurveyID;constraint:OnDelete:CASCADE" json:"questions"`
}

type SurveyQuestion struct {
	Model
	Question string `gorm:"not null;unique;size:512" json:"question"`
	SurveyID uint   `json:"surveyID"`
}

type BlacklistedTokens struct {
	Model
	Token string `gorm:"not null;unique"`
}
