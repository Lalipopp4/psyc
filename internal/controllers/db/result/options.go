package result

import (
	"context"
	"fmt"
	"psyc/internal/models"
	"strings"
)

func (r *resultRepository) GetUsers(ctx context.Context, key, param string) ([]string, error) {
	rows, err := r.cur.QueryContext(ctx, fmt.Sprintf("SELECT user_id FROM info WHERE %s=$1", key), param)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	users := make([]string, 100)
	var id string
	for rows.Next() {
		rows.Scan(&id)
		users = append(users, id)
	}
	return users, nil
}

func (r *resultRepository) GetByTest(ctx context.Context, test string, params ...string) ([]models.Res, error) {
	var clause string
	if len(params) > 0 {
		clause = fmt.Sprintf(" WHERE user_id IN (%s)", strings.Join(params, ", "))
	}
	rows, err := r.cur.QueryContext(ctx, fmt.Sprintf("SELECT res, user_id, date FROM %s%s", test, clause))
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]models.Res, 100)
	var res, userid, date string
	for rows.Next() {
		rows.Scan(&res, &date, &userid)
		results = append(results, models.Res{Res: res, Date: date, UserID: userid})
	}
	return results, nil
}

func (r *resultRepository) GetByUsers(ctx context.Context, users []string) ([]models.Test, error) {

	var results = make([]models.Test, len(r.tests))
	for i, t := range r.tests {
		var err error
		results[i].Results, err = r.GetByTest(ctx, t, users...)
		if err != nil {
			return nil, err
		}
		results[i].Test = t
	}
	return results, nil

}

func (r *resultRepository) Add(ctx context.Context, test *models.Test) error {
	_, err := r.cur.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s ('res', 'user_id', 'date') VALUES ($1, $2, $3);", test.Test), test.Results[0].Res, test.Results[0].UserID, test.Results[0].Date)
	return err
}
