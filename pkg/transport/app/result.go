package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"psyc/internal/models"
	"time"
)

func (rh *resultHandler) results(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var (
		results []models.Result
		err     error
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	switch {
	case query.Get("student") != "":
		results, err = rh.service.GetResults(ctx, &models.User{ID: query.Get("student_id")})
	case query.Get("group") != "":
		results, err = rh.service.GetResults(ctx, &models.User{Group: query.Get("group")})
	case query.Get("uni") != "":
		results, err = rh.service.GetResults(ctx, &models.User{Uni: query.Get("uni")})
	default:
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		rh.logger.Error(err)
		http.Redirect(w, r, "/index", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func (rh *resultHandler) addResult(w http.ResponseWriter, r *http.Request) {
	test := &models.Test{}
	if err := json.NewDecoder(r.Body).Decode(test); err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	result, err := rh.service.AddResult(ctx, test)
	if err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Result: %s", result)))
}

func (rh *resultHandler) getTest(w http.ResponseWriter, r *http.Request) {
	test := &models.Test{Test: r.URL.Query().Get("test")}
	if test.Test == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	questions, err := rh.service.GetTest(ctx, test)
	if err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func (rh *resultHandler) addReview(w http.ResponseWriter, r *http.Request) {
	review := &models.Review{}
	if err := json.NewDecoder(r.Body).Decode(review); err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := rh.service.AddReview(ctx, review); err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rh *resultHandler) getReview(w http.ResponseWriter, r *http.Request) {
	review := &models.Review{ResultID: r.URL.Query().Get("result_id")}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var err error
	review.Review, err = rh.service.GetReview(ctx, review)
	if err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(review)
}

func (rh *resultHandler) addTest(w http.ResponseWriter, r *http.Request) {
	test := &models.Test{}
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := rh.service.AddTest(ctx, test); err != nil {
		rh.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
