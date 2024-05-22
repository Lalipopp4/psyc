package result

import (
	"context"
	"fmt"
	"log"
	"psyc/internal/models"
)

func (r *resultRepository) addTest(ctx context.Context, test *models.Test) error {
	for _, q := range test.Questions {
		_, err := r.cur.ExecContext(ctx, "INSERT INTO test (name, code, type, question, answers, right_answer) VALUES ($1, $2, $3, $4, $5, $6);", test.Test, q.Code, q.Type, q.Question, q.Answers, q.RightAnswer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *resultRepository) Add(ctx context.Context, data any) error {
	var err error
	switch v := data.(type) {
	case *models.Test:
		return r.addTest(ctx, v)
	case *models.Result:
		_, err = r.cur.ExecContext(ctx, "INSERT INTO result (test, res, user_id, date) VALUES ($1, $2, $3, $4);", v.Test, v.Res, v.UserID, v.Date)
	case *models.Review:
		_, err = r.cur.ExecContext(ctx, "INSERT INTO review (user_id, review, author_id, result_id) VALUES ($1, $2, $3, $4);", v.UserID, v.Review, v.AuthorID, v.ResultID)
		if err != nil {
			return err
		}
		err = r.cache.AddReview(ctx, v)
	default:
		return fmt.Errorf("unknown type %T", v)
	}

	return err
}

func (r *resultRepository) getTest(ctx context.Context, test *models.Test) (map[uint]*models.Question, error) {
	rows, err := r.cur.QueryContext(ctx, "SELECT code, question, answers, right_answer FROM test WHERE name=$1", test.Test)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		code                           uint
		question, answers, rightAnswer string
		questions                      = make(map[uint]*models.Question)
	)
	for rows.Next() {
		rows.Scan(&code, &question, &answers, &rightAnswer)
		log.Println(code, "#", question, answers, rightAnswer)
		questions[code] = &models.Question{Question: question, Answers: answers, RightAnswer: rightAnswer}
	}

	return questions, nil
}

func (r *resultRepository) getReview(ctx context.Context, review *models.Review) (string, error) {
	rv, err := r.cache.GetReview(ctx, review)
	if err == nil {
		return rv, nil
	}
	row := r.cur.QueryRowContext(ctx, "SELECT user_id, review, author_id FROM review WHERE result_id=$1", review.ResultID)
	var userID, authorID, data string
	row.Scan(&data, &userID, &authorID)
	review.AuthorID = authorID
	review.UserID = userID
	review.Review = data

	return review.Review, nil
}

func (r *resultRepository) Get(ctx context.Context, arg any) (any, error) {
	switch v := arg.(type) {
	case *models.Review:
		return r.getReview(ctx, v)
	case *models.Test:
		return r.getTest(ctx, v)
	case *models.User:
		return r.getResults(ctx, v)
	}

	return nil, fmt.Errorf("unknown type")
}

func (r *resultRepository) getResults(ctx context.Context, user *models.User) ([]models.Result, error) {
	var (
		results = make([]models.Result, 100)
		query   string
	)
	switch {
	case user.ID != "":
		query = fmt.Sprintf("SELECT test, res, date FROM result WHERE user_id=%s", user.ID)
	case user.Email != "":
		query = fmt.Sprintf("SELECT test, res, date FROM result WHERE email=%s", user.Email)
	case user.Uni != "":
		query = fmt.Sprintf("SELECT test, res, date FROM result WHERE uni=%s", user.Uni)
	default:
		return nil, fmt.Errorf("unknown type")
	}
	rows, err := r.cur.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var test, res, date string
	for rows.Next() {
		rows.Scan(&test, &res, &date)
		results = append(results, models.Result{Test: test, Res: res, Date: date})
	}

	return results, nil
}
