package result

import (
	"context"
	"fmt"
	"psyc/internal/models"

	"golang.org/x/sync/errgroup"
)

var (
	columns = map[string]string{
		"keirsey": "user_id, res, type, date",
		"hall":    "user_id, EA, ME, SM, EM, ER, date",
	}
)

func (r *resultRepository) Get(ctx context.Context, param interface{}) ([]models.Result, error) {
	var wg errgroup.Group
	var types = map[string][]models.Result{
		"keirsey": make([]*models.TestKeirsey, 50),
	}
	for _, t := range r.tests {
		wg.Go(func() error {
			rows, err := r.cur.QueryContext(ctx, fmt.Sprintf("SELECT res, date FROM %s WHERE user_id=$1", t), param)
			if err != nil {
				return err
			}
			for rows.Next() {
				rows.Scan()
			}
			return nil
		})

	}

}

func (r *resultRepository) Add(ctx context.Context, res models.Result) error {
	var table, clmns, vals string
	switch res.(type) {
	case *models.TestKeirsey:
		table = "keirsey"
		clmns = columns["keirsey"]
		vals = "$1, $2, $3, $4"
	}
	_, err := r.cur.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, clmns, vals), res.Get()...)
	return err
}
