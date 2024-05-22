package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"psyc/internal/models"
)

func (a *userHandler) auth(w http.ResponseWriter, r *http.Request) {
	userAuth := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(userAuth); err != nil {
		a.logger.Error(err)
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	token, user, err := a.service.Login(r.Context(), userAuth.Email, userAuth.Password)
	if err != nil {
		a.logger.Error(err.Error())
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	if err := json.NewEncoder(w).Encode(token); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, user, http.StatusSeeOther)
}

func (a *userHandler) register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		a.logger.Error(err)
		http.Redirect(w, r, "/reg", http.StatusSeeOther)
		return
	}
	token, err := a.service.Register(r.Context(), user)
	if err != nil {
		a.logger.Error(err)
		http.Redirect(w, r, "/reg", http.StatusSeeOther)
		return
	}
	fmt.Println(token)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		a.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/auth", http.StatusSeeOther)
}

func (a *userHandler) info(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := a.service.Update(r.Context(), user); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
