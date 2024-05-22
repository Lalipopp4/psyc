package result

import (
	"context"
	"fmt"
	"log"
	"psyc/internal/models"
	"strconv"
	"time"
)

func (rs *resultService) AddResult(ctx context.Context, result *models.Test) (string, error) {
	return rs.tests[result.Test](ctx, result)
}

func (s *resultService) keirsey(ctx context.Context, result *models.Test) (string, error) {
	var (
		pairs          = [8]string{"E", "I", "S", "N", "T", "F", "J", "P"}
		profile, t, tp = "", 0, ""
		resBuf         = [4]int{0, 0, 0, 0}
	)
	for c, res := range result.Result {
		r, err := strconv.Atoi(res)
		if err != nil {
			return "", err
		}
		resBuf[int(c)%len(pairs)-1] = r
	}
	if t < 20 {
		tp = "Неяркий"
	} else {
		tp = "Яркий"
	}

	keirsey := &models.Result{Test: "keirsey", Res: fmt.Sprintf("%s, %s", profile, tp), UserID: result.UserID, Date: time.Now().String()}

	return keirsey.Res, s.repo.Add(ctx, keirsey)
}

func (s *resultService) hall(ctx context.Context, result *models.Test) (string, error) {
	var results = [5]int{}
	for c, res := range result.Result {
		r, err := strconv.Atoi(res)
		if err != nil {
			return "", err
		}
		results[(c-1)%5] += r
	}

	hall := &models.Result{
		Test: "hall", Res: fmt.Sprintf("СМК: %d, СММ: %d, СМЗ: %d, КМН: %d, ЭМП: %d",
			results[0], results[1], results[2], results[3], results[4]),
		UserID: result.UserID,
		Date:   time.Now().String(),
	}

	return hall.Res, s.repo.Add(ctx, hall)
}

func (s *resultService) bass(ctx context.Context, result *models.Test) (string, error) {
	var results = [3]int{}
	for _, res := range result.Result {
		index, err := strconv.Atoi(res)
		if err != nil {
			return "", err
		}
		results[index]++
	}
	bass := &models.Result{
		Test:   "bass",
		Res:    fmt.Sprintf("на себя: %d, на задачу: %d, на взаимодействие: %d", results[0], results[1], results[2]),
		UserID: result.UserID, Date: time.Now().String(),
	}

	return bass.Res, s.repo.Add(ctx, bass)
}

func (s *resultService) eysenck(ctx context.Context, result *models.Test) (string, error) {
	data, err := s.repo.Get(ctx, result.Test)
	if err != nil {
		return "", err
	}

	test := data.(map[uint]*models.Question)

	var res int
	for code, answer := range result.Result {
		if test[code].RightAnswer == answer {
			res++
		}
	}
	eysenck := &models.Result{Test: "eysenck", Res: strconv.Itoa(res), UserID: result.UserID, Date: time.Now().String()}
	return eysenck.Res, s.repo.Add(ctx, eysenck)
}

func (s *resultService) GetTest(ctx context.Context, test *models.Test) (map[uint]*models.Question, error) {
	data, err := s.repo.Get(ctx, test)
	if err != nil {
		return nil, err
	}

	return data.(map[uint]*models.Question), nil
}

func (s *resultService) AddReview(ctx context.Context, review *models.Review) error {
	return s.repo.Add(ctx, review)
}

func (s *resultService) AddTest(ctx context.Context, test *models.Test) error {
	return s.repo.Add(ctx, test)
}

func (s *resultService) GetReview(ctx context.Context, review *models.Review) (string, error) {
	data, err := s.repo.Get(ctx, review)
	if err != nil {
		return "", err
	}
	log.Println(review)

	return data.(string), nil
}

func (s *resultService) GetResults(ctx context.Context, user *models.User) ([]models.Result, error) {
	data, err := s.repo.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	return data.([]models.Result), nil
}
