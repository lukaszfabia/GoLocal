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
	FirstName    string     `gorm:"not null:size:255" json:"firstName"`
	LastName     string     `gorm:"not null:size:255" json:"lastName"`
	Email        string     `gorm:"not null;size:100;unique" json:"email"`
	Password     *string    `gorm:"null;size:400" json:"-"`
	Birthday     *time.Time `gorm:"type:date"`
	Bio          *string    `gorm:"size:512" json:"bio"`
	AuthProvider *string    `gorm:"null" json:"provider"`
	IsVerified   bool       `gorm:"null" json:"isVerified"`

	IsPremium bool    `gorm:"default:false" json:"isPremium"`
	AvatarURL *string `gorm:"null;size:1024" json:"avatarUrl"`

	Followers []*User `gorm:"many2many:user_followers;constraint:OnDelete:CASCADE" json:"followers"`
	Following []*User `gorm:"many2many:user_following;constraint:OnDelete:CASCADE" json:"following"`

	Comments []*Comment    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"comments"`
	Votes    []*VoteAnswer `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"votes"`

	Location   *Location      `gorm:"constraint:OnDelete:CASCADE" json:"location"`
	LocationID *uint          `json:"locationID"`
	Devices    []*DeviceToken `gorm:"many2many:devices;constraint:OnDelete:CASCADE" json:"devices"`

	SkipValidation bool `gorm:"-" json:"-"`
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
	Geom      string  `gorm:"column:geom;type:geometry(Point,4326)" json:"-"`
}

type Address struct {
	Model
	Street         string `gorm:"not null;size:255" json:"street"`
	StreetNumber   string `gorm:"null" json:"streetNumber"`
	AdditionalInfo string `gorm:"null:size:512" json:"additionalInfo"`
}

type Event struct {
	Model
	EventOrganizers []*User `gorm:"many2many:event_organizers" json:"eventOrganizers"`

	Title       string    `gorm:"not null;size:255" json:"title"`
	Description string    `gorm:"default:'';size:255" json:"description"`
	ImageURL    *string   `gorm:"size:1024" json:"image"`
	IsAdultOnly bool      `gorm:"default:true" json:"isAdultOnly"`
	EventType   EventType `gorm:"type:text;not null" json:"event_type"`
	Tags        []*Tag    `gorm:"many2many:event_tags" json:"event_tags"` // for ml

	// Timestamp with time zone
	StartDate  *time.Time `gorm:"type:date;not null" json:"startDate"`
	FinishDate *time.Time `gorm:"type:date;" json:"finishDate"`

	Location   *Location `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE" json:"location"`
	LocationID uint      `json:"locationID"`

	Comments []*Comment `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"comments"`
	Votes    []*Vote    `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"votes"`
}

type Tag struct {
	Model
	Name string `gorm:"not null;unique;size:100"`
}

type Comment struct {
	Model
	UserID  uint   `json:"userID"`
	EventID uint   `json:"eventID"`
	Content string `gorm:"not null;size:255" json:"content"`
}

type Vote struct {
	Model
	Event    Event        `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:RESTRICT" json:"event"`
	EventID  uint         `json:"eventID"`
	Text     string       `gorm:"not null;size:255" json:"text"`
	VoteType VoteType     `gorm:"not null" json:"voteType"`
	Options  []VoteOption `gorm:"foreignKey:VoteID" json:"options"`
	EndDate  *time.Time   `gorm:"type:date" json:"endDate"`
}

type VoteAnswer struct {
	Model
	VoteID       uint       `json:"voteId"`
	UserID       uint       `json:"userId"`
	VoteOptionID uint       `json:"optionSelectedId"`
	VoteOption   VoteOption `json:"optionSelected"`
}

type VoteOption struct {
	Model
	Vote                Vote                `gorm:"foreignKey:VoteID;references:ID;constraint:OnDelete:RESTRICT" json:"vote"`
	VoteID              uint                `json:"voteId"`
	Text                string              `gorm:"not null;size:255" json:"text"`
	ParticipationStatus ParticipationStatus `gorm:"not null" json:"participationStatus"`
	VoteAnswers         []VoteAnswer        `gorm:"foreignKey:VoteOptionID" json:"voteAnswers"`
}

type Opinion struct {
	Model
	Event   *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:RESTRICT" json:"event"`
	EventID uint   `json:"eventID"`

	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:RESTRICT" json:"user"`
	UserID uint  `json:"userID"`

	Opinion string `gorm:"not null;size:1024"`
}

type BlacklistedTokens struct {
	Model
	Token string `gorm:"not null;unique"`
}

type PreferenceSurvey struct {
	Model
	Title       string                     `gorm:"not null;size:255" json:"title"`
	Description string                     `gorm:"size:1024" json:"description"`
	Questions   []PreferenceSurveyQuestion `gorm:"foreignKey:SurveyID" json:"questions"`
}

type PreferenceSurveyQuestion struct {
	Model
	Survey   PreferenceSurvey         `gorm:"foreignKey:SurveyID" json:"survey"`
	SurveyID uint                     `json:"surveyID"`
	Text     string                   `gorm:"not null;size:1024" json:"text"`
	Type     QuestionType             `gorm:"not null" json:"type"`
	Options  []PreferenceSurveyOption `gorm:"foreignKey:QuestionID" json:"options"`
}

type PreferenceSurveyOption struct {
	Model
	Question    PreferenceSurveyQuestion `gorm:"foreignKey:QuestionID" json:"question"`
	QuestionID  uint                     `json:"questionID"`
	Text        string                   `gorm:"not null;size:1024" json:"text"`
	Tag         Tag                      `gorm:"foreignKey:TagID" json:"tag"`
	TagID       uint                     `json:"tagID"`
	TagPositive bool                     `json:"tagPositive"`
}

type PreferenceSurveyAnswer struct {
	Model
	Survey          PreferenceSurvey               `gorm:"foreignKey:SurveyID" json:"survey"`
	SurveyID        uint                           `json:"surveyID"`
	Question        PreferenceSurveyQuestion       `gorm:"foreignKey:QuestionID" json:"question"`
	QuestionID      uint                           `json:"questionID"`
	User            User                           `gorm:"foreignKey:UserID" json:"user"`
	UserID          uint                           `json:"userID"`
	SelectedOptions []PreferenceSurveyAnswerOption `gorm:"foreignKey:AnswerID" json:"options"`
}

type PreferenceSurveyAnswerOption struct {
	Model
	Answer   PreferenceSurveyAnswer `gorm:"foreignKey:AnswerID" json:"answer"`
	AnswerID uint                   `json:"answerID"`
	Option   PreferenceSurveyOption `gorm:"foreignKey:OptionID" json:"option"`
	OptionID uint                   `json:"optionID"`
}

type UserPreference struct {
	Model
	User   User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	UserID uint  `json:"userID"`
	Tags   []Tag `gorm:"many2many:user_preferences_tags" json:"tags"`
}

type DeviceToken struct {
	Model
	// FMC token
	Token     string  `gorm:"not null;size:1024" json:"token"`
	OSVersion *string `gorm:"size:32" json:"os"`
	Platform  *string `gorm:"size:32" json:"platform"`
	Users     []*User `gorm:"many2many:user_devices;" json:"users"`
}

type ErrorResponse struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}
