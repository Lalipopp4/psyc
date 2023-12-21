package transport

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

func (a *appHTTP) hall(w http.ResponseWriter, r *http.Request) {
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
	err := a.result.Hall(r.Context(), r.Context().Value("id").(string), res)
	if err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func (a *appHTTP) keirsey(w http.ResponseWriter, r *http.Request) {
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
	if err := a.result.Keirsey(r.Context(), r.Context().Value("id").(string), res); err != nil {
		a.logger.Error("error handling results: %v", err)
		http.Redirect(w, r, "/user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
