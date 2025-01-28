package forms // aka dto

import (
	"backend/internal/models"
	"mime/multipart"
	"time"
)

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RefreshTokenRequest struct {
	Token string `json:"refresh"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditAccount struct {
	FirstName string                `json:"firstName" form:"firstName"`
	LastName  string                `json:"lastName" form:"lastName"`
	Email     string                `json:"email" form:"email"`
	Password  string                `json:"password" form:"password"`
	Birthday  string                `json:"birthday" form:"birthday"` // 1970-01-01
	Bio       string                `json:"bio" form:"bio"`
	Avatar    *multipart.FileHeader `json:"-" form:"avatar"`
}

type RestoreAccount struct {
	Email string `json:"email"`
}

type CodeRequest struct {
	Code string `json:"code"`
}

type NewPasswordRequest struct {
	Password string `json:"newPassword"`
}

type Device struct {
	UserID    int     `json:"userID"`
	Token     string  `json:"token"`
	OSVersion *string `json:"os"`
	Platform  *string `json:"platform"`
}

type Event struct {
	// list with users ids
	Organizers  []uint                `json:"organizers" form:"organizers"`
	Title       string                `json:"title" form:"title"`
	Description string                `json:"description" form:"description"`
	Image       *multipart.FileHeader `json:"-" form:"image"`
	StartDate   *time.Time            `json:"startDate" form:"startDate"`
	FinishDate  *time.Time            `json:"finishDate" form:"finishDate"`
	IsAdultOnly bool                  `form:"isAdultOnly" json:"isAdultOnly"`
	EventType   string                `json:"eventType" form:"eventType"`
	ImageURL    string                `json:"-" form:"-"`
	Tags        []string              `json:"tags" form:"tags"` // list with user input tags

	Lon string `json:"lon" form:"lon"` // coords to get more info
	Lat string `json:"lat" form:"lat"`
}

type VoteInVotingForm struct {
	VoteID       int `json:"voteID"`
	VoteOptionID int `json:"voteOptionID"`
}

type VoteForm struct {
	ID       int              `json:"id"`
	EventID  int              `json:"eventID"`
	Text     string           `json:"text"`
	VoteType string           `json:"voteType"`
	EndDate  time.Time        `json:"endDate"`
	Options  []VoteOptionForm `json:"options"`
	Event    models.Event     `json:"event"`
}

type VoteOptionForm struct {
	ID         int    `json:"id"`
	VoteID     int    `json:"voteID"`
	Text       string `json:"text"`
	IsSelected bool   `json:"isSelected"`
	VotesCount int    `json:"votesCount"`
}

type PreferenceSurveyAnswer struct {
	PreferenceSurveyID int   `json:"SurveyID"`
	QuestionID         int   `json:"QuestionID"`
	UserID             int   `json:"UserID"`
	Options            []int `json:"Options"`
}

type ReportForm struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}
