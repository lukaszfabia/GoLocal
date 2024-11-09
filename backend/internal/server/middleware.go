package server

import (
	"backend/internal/auth"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
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

		if s.db.TokenService().IsTokenBlacklisted(tokenStr) {
			log.Println("Token is blacklisted")
			s.NewResponse(w, http.StatusUnauthorized, "")
			return
		}

		// decode token
		id, err := auth.DecodeJWT(tokenStr)

		if err != nil {
			log.Println("Error during decoding token")
			s.NewResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := s.db.UserService().GetUser(fmt.Sprintf("id = %d", id))

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
	})
}
