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

	Location   *Location `gorm:"constraint:OnDelete:CASCADE" json:"location"`
	LocationID *uint     `json:"locationID"`

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
	EventID        uint           `json:"eventID"`
	Text           string         `gorm:"not null;size:255" json:"text"`
	VoteType       VoteType       `gorm:"not null" json:"voteType"`
	VoteChangeType VoteChangeType `gorm:"not null" json:"voteChangeType"`
	Options        []VoteOption   `gorm:"foreignKey:QuestionID" json:"options"`
}

type VoteAnswer struct {
	gorm.Model
	Vote       Vote               `json:"vote"`
	VoteID     uint               `json:"voteId"`
	QuestionID uint               `json:"questionId"`
	UserID     uint               `json:"userId"`
	Options    []VoteAnswerOption `gorm:"foreignKey:AnswerID" json:"options"`
}

type VoteOption struct {
	gorm.Model
	VoteQuestion Vote   `gorm:"foreignKey:QuestionID" json:"voteQuestion"`
	QuestionID   uint   `json:"questionId"`
	Text         string `gorm:"not null;size:255" json:"text"`
}

type VoteAnswerOption struct {
	gorm.Model
	AnswerID uint `json:"answerID"`
	OptionID uint `json:"optionID"`
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
	gorm.Model
	Title       string                     `gorm:"not null;size:255" json:"title"`
	Description string                     `gorm:"size:1024" json:"description"`
	Questions   []PreferenceSurveyQuestion `gorm:"foreignKey:SurveyID" json:"questions"`
}

type PreferenceSurveyQuestion struct {
	gorm.Model
	SurveyID uint                     `json:"surveyID"`
	Text     string                   `gorm:"not null;size:1024" json:"text"`
	Type     QuestionType             `gorm:"not null" json:"type"`
	Options  []PreferenceSurveyOption `gorm:"foreignKey:QuestionID" json:"options"`
	Toggle   *bool                    `json:"toggle"` // Only used if Type == Toggle
}

// PreferenceSurveyOption represents an option for SingleChoice or MultipleChoice
type PreferenceSurveyOption struct {
	gorm.Model
	QuestionID uint   `json:"questionID"`
	Text       string `gorm:"not null;size:1024" json:"text"`
	IsSelected bool   `json:"isSelected"` // Used for MultipleChoice answers
}

type PreferenceSurveyAnswer struct {
	gorm.Model
	SurveyID   uint                           `json:"surveyID"`
	QuestionID uint                           `json:"questionID"`
	UserID     uint                           `json:"userID"`
	Toggle     *bool                          `json:"toggle"`                             // For Toggle type
	OptionID   *uint                          `json:"optionID"`                           // For SingleChoice
	Options    []PreferenceSurveyAnswerOption `gorm:"foreignKey:AnswerID" json:"options"` // For MultipleChoice, store option IDs
}

type PreferenceSurveyAnswerOption struct {
	gorm.Model
	AnswerID uint `json:"answerID"`
	OptionID uint `json:"optionID"`
}

type Recommendation struct {
	Model
	UserID uint   `json:"userID"`
	Text   string `gorm:"not null;size:1024" json:"text"`
}
