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

func (a *AccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.get(w, r)

	case http.MethodDelete:
		a.delete(w, r)

	case http.MethodPut, http.MethodPatch:
		put(w, r)

	default:
		app.NewResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (a *AccountHandler) post(w http.ResponseWriter, r *http.Request) {}
func (a *AccountHandler) get(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		app.NewResponse(w, http.StatusUnauthorized, nil)
		return
	}

	app.NewResponse(w, http.StatusOK, user)
}
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
		// Handle image
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
