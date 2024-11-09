package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

/*
Generates token with given life time
Decoding token returns id of user

Params:

  - duration time.Duration: token life time
  - message(data)

Returns:

  - generated token as a string
*/
func generate(duration time.Duration, id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(duration).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return ""
	}

	return tokenStr
}

/*
Generate refresh and access tokens, no need to check tokens are empty

Params:

  - id uint: user id
  - additional block with shoud be executed

Returns:

  - tokens *Token: struct with both tokens
  - error occured during operation
*/
func GenerateJWT(id uint, additionalFunc *func() error) (*Token, error) {
	if additionalFunc != nil {
		if err := (*additionalFunc)(); err != nil {
			return nil, err
		}
	}

	// duration 24h
	access := generate(time.Hour*24, id)

	// duration 30 days
	refresh := generate(time.Hour*30*24, id)

	if access != "" && refresh != "" {
		return &Token{
			Access:  access,
			Refresh: refresh,
		}, nil
	}

	log.Println("Something went wrong during token generating")
	return nil, errors.New("failed to generate tokens!")
}

/*
Decodes token

Params:

  - tokenStr string: token received in header
  - serivce database.Service: checks is token blacklisted

Returns:

  - id uint: decoded sub
  - error occured during executing
*/
func DecodeJWT(tokenStr string) (uint, error) {

	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		msg := "Error during paring the token"
		log.Println(msg)
		return 0, errors.New(strings.ToLower(msg))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok || time.Now().Unix() > int64(exp) {
			msg := "Token expired or invalid"
			log.Println(msg)
			return 0, errors.New(strings.ToLower(msg))
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			msg := "Invalid token subject"
			log.Println(msg)
			return 0, errors.New(strings.ToLower(msg))
		} else {
			return uint(sub), nil
		}

	} else {
		msg := "Invalid token claims"
		log.Println(msg)
		return 0, errors.New(strings.ToLower(msg))
	}
}
