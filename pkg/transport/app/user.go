package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"psyc/internal/models"
)

func (a *userHandler) auth(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
	)
	token, user, err := a.service.Login(r.Context(), email, password)
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
	user := &models.User{
		Password: r.FormValue("password"),
		Info: models.Info{
			Lastname:   r.FormValue("lastname"),
			Firstname:  r.FormValue("firstname"),
			Patronymic: r.FormValue("patronymic"),
			Email:      r.FormValue("email"),
			Uni:        r.FormValue("uni"),
			Age:        r.FormValue("age"),
			Grade:      r.FormValue("grade"),
			Syllabus:   r.FormValue("syllabus"),
			City:       r.FormValue("city"),
		},
	}
	fmt.Println(user)
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
	user := &models.User{
		ID:       r.Context().Value("id").(string),
		Password: r.FormValue("password"),
		Info: models.Info{
			Lastname:   r.FormValue("lastname"),
			Firstname:  r.FormValue("firstname"),
			Patronymic: r.FormValue("patronymic"),
			Email:      r.FormValue("email"),
			Uni:        r.FormValue("uni"),
			Age:        r.FormValue("age"),
			Grade:      r.FormValue("grade"),
			Syllabus:   r.FormValue("syllabus"),
			City:       r.FormValue("city"),
		},
	}
	if err := a.service.Update(r.Context(), user); err != nil {
		a.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
