package account

import (
	"backend/internal/app"
	"backend/internal/database"
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/image"
	"backend/pkg/parsers"
	"log"
	"net/http"
)

type AccountHandler struct {
	UserService database.UserService
}

// @Summary Handle account actions
// @Description Handles requests for account management
// @Tags account
// @Accept json, multipart/form-data
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Failure 405 {object} map[string]string
func (a *AccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.get(w, r)
		return
	case http.MethodDelete:
		a.delete(w, r)
		return
	case http.MethodPut, http.MethodPatch:
		a.put(w, r)
		return
	default:
		app.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// @Summary Get user account
// @Description Retrieve current authenticated user data
// @Tags account
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Router /api/auth/account/ [get]
func (a *AccountHandler) get(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		app.NewResponse(w, http.StatusUnauthorized, nil)
		return
	}

	app.NewResponse(w, http.StatusOK, user)
}

// @Summary Delete user account
// @Description Delete the authenticated user account
// @Tags account
// @Accept json
// @Produce json
// @Success 200 {string} string "User deleted successfully"
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/auth/account/ [delete]
func (a *AccountHandler) delete(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		app.NewResponse(w, http.StatusUnauthorized, nil)
		return
	}

	if err := a.UserService.DeleteUser(user); err != nil {
		app.NewResponse(w, http.StatusNotFound, "User not found")
		return
	}
	app.NewResponse(w, http.StatusOK, "User deleted successfully")
}

// @Summary Update user account
// @Description Update authenticated user account details
// @Tags account
// @Accept json, multipart/form-data
// @Produce json
// @Param account body forms.EditAccount true "Updated account data"
// @Param avatar formData file false "New avatar image"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/account/ [put]
func (a *AccountHandler) put(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		app.NewResponse(w, http.StatusUnauthorized, nil)
		return
	}

	form, ok := r.Context().Value(_userForm).(*forms.EditAccount)
	if !ok {
		app.InvalidDataResponse(w)
		return
	}

	info, err := parsers.GetFileFromForm(r.MultipartForm, "avatar")
	if err != nil {
		log.Println("Avatar upload error:", err)
	} else {
		info.OldPath = user.AvatarURL
		if path, err := image.SaveImage[*image.Avatar](info); err == nil {
			user.AvatarURL = &path
		} else {
			log.Println("Failed to save avatar:", err)
		}
	}

	updatedUser, err := a.UserService.SaveUser(form, user)
	if err != nil {
		log.Println("User save error:", err)
		app.NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	app.NewResponse(w, http.StatusOK, updatedUser)
}
