package models

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Validation

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password is empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashed), nil
}

// Checks is len of string is in [min, max]
func isBetween(min, max int, str string) bool {
	l := len(str)

	return l >= min && l <= max
}

func isSecurePassword(password string) bool {
	const midReg string = `^[A-Za-z\d]{5,}$`           // one big letter & one decimal & len = 8
	const strongReg string = `^[A-Za-z\d@$!%#?&]{5,}$` // the same above + special char

	goodPasswordValidator := regexp.MustCompile(midReg)
	strongPasswordValidator := regexp.MustCompile(strongReg)

	return (goodPasswordValidator.MatchString(password) ||
		strongPasswordValidator.MatchString(password))
}

func isHashed(password string) bool {
	return strings.HasPrefix(password, "$2a$") || strings.HasPrefix(password, "$2b$")
}

// Hook for validation for really important data
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.SkipValidation {
		return nil
	}

	if !isBetween(1, 255, u.FirstName) || !isBetween(1, 255, u.LastName) {
		return errors.New("first name and last name are required")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email format")
	}

	if u.Password != nil && !isHashed(*u.Password) {
		if !isSecurePassword(*u.Password) {
			return errors.New("password must be at least 5 characters")
		} else {
			hashedPassword, err := HashPassword(*u.Password)
			if err != nil {
				return err
			}
			u.Password = &hashedPassword
		}
	}

	return nil
}
