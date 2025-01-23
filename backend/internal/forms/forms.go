package forms // aka dto

import (
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
	// list with tags as ids
	Tags       []uint `json:"tags" form:"tags"`
	LocationID uint   `json:"locationID" form:"locationID"`
	ImageURL   string `json:"-" form:"-"`
}
