package forms

import (
	"mime/multipart"
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
