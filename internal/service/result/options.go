package result

import (
	"context"
	"fmt"
	"psyc/internal/models"
	"strconv"
	"time"
)

func (s *resultService) Get(ctx context.Context, key string, param string) ([]models.Test, error) {
	switch key {
	case "table":
		test, err := s.repo.GetByTest(ctx, param)
		if err != nil {
			return nil, err
		}
		return []models.Test{{Test: param, Results: test}}, nil
	case "user":
		return s.repo.GetByUsers(ctx, []string{param})
	default:
		users, err := s.repo.GetUsers(ctx, key, param)
		if err != nil {
			return nil, err
		}
		return s.repo.GetByUsers(ctx, users)
	}
}

func (s *resultService) Keirsey(ctx context.Context, id string, res [4]int) error {
	var (
		pairs          = [8]string{"E", "I", "S", "N", "T", "F", "J", "P"}
		profile, t, tp = "", 0, ""
	)
	for i := 0; i < 4; i++ {
		switch {
		case res[i] > 4:
			profile += pairs[i*2]
		default:
			profile += pairs[i*2+1]
		}
		if i == 0 {
			t += (res[i] - 5) * 2
		} else {
			t += res[i] - 10
		}
	}
	if t < 20 {
		tp = "Неяркий"
	} else {
		tp = "Яркий"
	}
	keirsey := models.Res{Res: fmt.Sprintf("%s, %s", profile, tp), UserID: id, Date: time.Now().String()}
	test := &models.Test{Test: "keirsey", Results: []models.Res{keirsey}}
	return s.repo.Add(ctx, test)
}

func (s *resultService) Hall(ctx context.Context, id string, res [5]int) error {
	temp := fmt.Sprintf("СМК: %d, СММ: %d, СМЗ: %d, КМН: %d, ЭМП: %d", res[0], res[1], res[2], res[3], res[4])
	hall := models.Res{Res: temp, UserID: id, Date: time.Now().String()}
	test := &models.Test{Test: "hall", Results: []models.Res{hall}}
	return s.repo.Add(ctx, test)
}

func (s *resultService) Bass(ctx context.Context, id string, self, task, social int) error {
	bass := models.Res{Res: fmt.Sprintf("на себя: %d, на задачу: %d, на взаимодействие: %d", self, task, social), UserID: id, Date: time.Now().String()}
	test := &models.Test{Test: "bass", Results: []models.Res{bass}}
	return s.repo.Add(ctx, test)
}

func (s *resultService) Eysenck(ctx context.Context, id string, res int) error {
	eysenck := models.Res{Res: strconv.Itoa(res), UserID: id, Date: time.Now().String()}
	test := &models.Test{Test: "eysenck", Results: []models.Res{eysenck}}
	return s.repo.Add(ctx, test)
}
