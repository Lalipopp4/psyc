package transport

import (
	"encoding/json"
	"net/http"

	"psyc/internal/models"
)

// func (a *appHTTP) account(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseFiles(`C:\Users\anton\Go\src\github.com\Lalipopp4\test_server\ui\templates\user.html`)
// 	if err != nil {
// 		a.logger.Error("template error: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	id := r.Context().Value("id").(string)
// 	results, err := a.result.Get(id)
// 	if err != nil {
// 		a.logger.Error("template error: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, results)

// 	info, _ := sessions.Store.Get(r, "psycProfiling")
// 	rows, err := DB.Query("SELECT res, type FROM test_keirsey WHERE user_id=$1", info.Values["id"].(int))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	var (
// 		data = Data{
// 			Name:    info.Values["name"].(string),
// 			Profile: []Res{},
// 		}
// 	)
// 	f, err := os.Open(`C:\Users\anton\Go\src\github.com\Lalipopp4\test_server\internal\answers\answers.txt`)
// 	if err != nil {
// 		log.Println(err)

// 	}
// 	defer f.Close()
// 	for i := 1; rows.Next(); i++ {
// 		r := Res{N: i}
// 		err := rows.Scan(&r.Formula, &r.Type)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		scanner := bufio.NewScanner(f)
// 		for scanner.Scan() {
// 			row := scanner.Text()
// 			if row[:4] == r.Formula {
// 				r.Description = row[7:]
// 				r.Pers = Pers[i/5]
// 				data.Profile = append(data.Profile, r)
// 			}
// 		}
// 	}
// 	templ.Execute(w, data)
// }

func (a *appHTTP) login(w http.ResponseWriter, r *http.Request) {
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
	)
	token, err := a.user.Login(r.Context(), email, password)
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
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *appHTTP) register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Info: models.Info{
			Lastname:   r.FormValue("lastname"),
			Firstname:  r.FormValue("firstname"),
			Patronymic: r.FormValue("patronymic"),
			Email:      r.FormValue("email"),
			Password:   r.FormValue("password"),
		},
	}
	token, err := a.user.Register(r.Context(), user)
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
