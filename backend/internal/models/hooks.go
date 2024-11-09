package models

import (
	"errors"
	"log"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
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

func isSecurePassword(password string, empty bool) bool {
	const midReg string = `^[A-Za-z\d]{8,}$`           // one big letter & one decimal & len = 8
	const strongReg string = `^[A-Za-z\d@$!%#?&]{8,}$` // the same above + special char

	goodPasswordValidator := regexp.MustCompile(midReg)
	strongPasswordValidator := regexp.MustCompile(strongReg)

	var basicCond bool = goodPasswordValidator.MatchString(password) || strongPasswordValidator.MatchString(password)

	if empty {
		return basicCond || password == ""
	}

	return basicCond
}

// Hook for validation and processing
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("Validating user data")
	if !isBetween(1, 255, u.FirstName) || !isBetween(1, 255, u.LastName) {
		return errors.New("first name and last name are required")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email format")
	}

	if !isSecurePassword(*u.Password, u.Password != nil) {
		return errors.New("password must be at least 8 characters")
	}

	if u.Password != nil {
		hashedPassword, err := hashPassword(*u.Password)
		if err != nil {
			return err
		}
		u.Password = &hashedPassword
	}

	return nil
}
