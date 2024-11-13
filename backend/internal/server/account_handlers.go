package server

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg"
	"log"
	"net/http"
)

// Manage user
func (s *Server) AccountHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from ctx
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "Unauthorized access")
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.NewResponse(w, http.StatusOK, user)

	case http.MethodDelete:
		if err := s.db.UserService().DeleteUser(user); err != nil {
			s.NewResponse(w, http.StatusNotFound, "User not found")
			return
		}
		s.NewResponse(w, http.StatusOK, "User deleted successfully")

	case http.MethodPut, http.MethodPatch:
		s.handleUpdateUser(w, r, user)

	default:
		s.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request, user *models.User) {
	form, err := pkg.DecodeMultipartForm[forms.EditAccount](r)
	if err != nil {
		s.InvalidFormResponse(w)
		return
	}

	info, err := pkg.GetFileFromForm(r.MultipartForm, "avatar")
	if err != nil {
		log.Println("Avatar upload error:", err)
	} else {
		// Handle image
		info.OldPath = user.AvatarURL
		if path, err := pkg.SaveImage[*pkg.Avatar](info); err == nil {
			user.AvatarURL = &path
		} else {
			log.Println("Failed to save avatar:", err)
		}
	}

	// Assign all form fields to the user model
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.Email = form.Email
	user.Password = &form.Password
	user.Bio = &form.Bio
	user.Birthday = &form.Birthday

	updatedUser, err := s.db.UserService().SaveUser(user)
	if err != nil {
		log.Println("User save error:", err)
		s.NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	s.NewResponse(w, http.StatusOK, updatedUser)
}
