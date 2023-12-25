package transport

import (
	"encoding/json"
	"net/http"

	"psyc/internal/models"
)

func (a *userHandler) login(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
	)
	token, user, err := a.service.Login(r.Context(), email, password)
	if err != nil {
		a.logger.Error(err.Error())
		http.Redirect(w, r, "/auth", http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(token); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, user, http.StatusSeeOther)
}

func (a *userHandler) register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Info: models.Info{
			Lastname:   r.FormValue("lastname"),
			Firstname:  r.FormValue("firstname"),
			Patronymic: r.FormValue("patronymic"),
			Email:      r.FormValue("email"),
			Password:   r.FormValue("password"),
		},
	}
	token, err := a.service.Register(r.Context(), user)
	if err != nil {
		a.logger.Error(err.Error())
		http.Redirect(w, r, "/reg", http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(token); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/auth", http.StatusSeeOther)
}
