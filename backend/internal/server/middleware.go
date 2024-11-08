package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) isAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		// no bearer token
		if authorization == "" {
			log.Println("Unauthorized")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}

		tokenStr := strings.TrimPrefix(authorization, "Bearer ")

		if tokenStr == "" {
			log.Println("Token missing after trimming Bearer prefix")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}

		// check is token blacklisted
		if s.db.TokenService().IsTokenBlacklisted(tokenStr) {
			log.Println("Token token has been expired")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}

		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil {
			log.Println("Error during paring the token")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp, ok := claims["exp"].(float64)
			if !ok || time.Now().Unix() > int64(exp) {
				log.Println("Token expired or invalid")
				s.NewResponse(w, http.StatusUnauthorized, "")
				return
			}

			sub, ok := claims["sub"].(float64)
			if !ok {
				log.Println("Invalid token subject")
				s.NewResponse(w, http.StatusUnauthorized, "")
				return
			}

			user, err := s.db.UserService().GetUser(fmt.Sprintf("id = %d", uint(sub)))

			if err != nil {
				log.Println("Error retrieving user from database:", err)
				s.NewResponse(w, http.StatusUnauthorized, "")
				return
			}

			// set user in ctx
			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)

			// go to next handler
			next.ServeHTTP(w, r)

		} else {
			log.Println("Invalid token claims")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}
	})
}
