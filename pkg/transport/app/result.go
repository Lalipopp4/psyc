package transport

import (
	"log"
	"net/http"
	"psyc/internal/models"
	"strconv"
	"sync"
	"text/template"
)

func (a *resultHandler) hall(w http.ResponseWriter, r *http.Request) {
	var (
		res = [5]int{0, 0, 0, 0, 0}
		wg  sync.WaitGroup
	)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 1; j < 7; j++ {
				val, err := strconv.Atoi(r.FormValue("q" + strconv.Itoa(i*6+j)))
				if err != nil {
					log.Println(err)
					return
				}
				res[i] += val
			}
		}(i)
	}
	wg.Wait()
	err := a.service.Hall(r.Context(), r.Context().Value("id").(string), res)
	if err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *resultHandler) keirsey(w http.ResponseWriter, r *http.Request) {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		res = [4]int{0, 0, 0, 0}
	)
	wg.Add(9)
	for i := 0; i < 9; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 1; j < 8; j++ {
				val, err := strconv.Atoi(r.FormValue("q" + strconv.Itoa(i*7+j)))
				if err != nil {
					a.logger.Error("error parsing results: %v", err)
					return
				}
				mu.Lock()
				res[j/2] += val
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	if err := a.service.Keirsey(r.Context(), r.Context().Value("id").(string), res); err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *resultHandler) bass(w http.ResponseWriter, r *http.Request) {
	res := [3]int{0, 0, 0}
	for i := 0; i < 27; i++ {
		val, err := strconv.Atoi(r.FormValue("q" + strconv.Itoa(i)))
		if err != nil {
			a.logger.Error("error parsing results: %v", err)
			return
		}
		res[val]++
	}
	if err := a.service.Bass(r.Context(), r.Context().Value("id").(string), res[0], res[1], res[2]); err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *resultHandler) eysenck(w http.ResponseWriter, r *http.Request) {
	res := 0
	for i := 1; i < 41; i++ {
		if eysenckResults[i-1] == r.FormValue("q"+strconv.Itoa(i)) {
			res++
		}
	}
	if err := a.service.Eysenck(r.Context(), r.Context().Value("id").(string), res); err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *resultHandler) account(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("psyc/static/templates/user.html"))
	results, err := a.service.Get(r.Context(), "user", r.Context().Value("id").(string))
	if err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	profile := struct {
		Name  string
		Tests []models.Test
	}{
		"Hello",
		results,
	}
	tmpl.Execute(w, profile)
}

func (a *resultHandler) admin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("psyc/static/templates/admin.html"))
	results, err := a.service.Get(r.Context(), r.FormValue("search"), r.FormValue("param"))
	if err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, results)
}
