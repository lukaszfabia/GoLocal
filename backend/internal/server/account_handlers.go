package server

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg"
	"log"
	"net/http"
)

func (s *Server) AccountHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		s.NewResponse(w, http.StatusUnauthorized, "")
		return
	}

	switch r.Method {
	case http.MethodGet:
		s.NewResponse(w, http.StatusOK, user)
	case http.MethodDelete:
		// delete acc
		if err := s.db.UserService().DeleteUser(user); err != nil {
			s.NewResponse(w, http.StatusNotFound, err.Error())
			return
		}
		s.NewResponse(w, http.StatusOK, "User deleted")

	case http.MethodPut:
	case http.MethodPatch:
		// update account
		form, err := pkg.DecodeMultipartForm[forms.EditAccount](r)

		if err != nil {
			s.InvalidFormResponse(w)
		}

		info, err := pkg.GetFileFromForm(r.MultipartForm, "avatar") // just dont update image if there was a problem
		if err != nil {
			log.Println(err)
			s.NewResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		info.OldPath = user.AvatarURL

		if path, err := pkg.SaveImage[*pkg.Avatar](info); err != nil {

		} else {
			//assign all
			user.FirstName = form.FirstName
			user.LastName = form.LastName
			user.Email = form.Email
			user.Password = &form.Password
			user.AvatarURL = &path
			user.Bio = &form.Bio
			user.Birthday = &form.Birthday

			if updated, err := s.db.UserService().SaveUser(user); err != nil {
				log.Println(err)
				s.NewResponse(w, http.StatusBadRequest, err.Error())
			} else {
				s.NewResponse(w, http.StatusOK, updated)
			}
		}

	default:
		s.NewResponse(w, http.StatusBadRequest, "")
	}

}
