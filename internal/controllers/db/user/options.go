package user

import (
	"context"
	"psyc/internal/models"
)

func (r *userRepository) Add(ctx context.Context, user *models.User) error {
	tx, err := r.cur.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into users (id, email) values ($1, $2)", user.ID, user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, `insert into info (user_id, email, last_name, first_name, patronymic, group, grade, sullabus, city) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, user.ID, user.Info.Email,
		user.Info.Lastname, user.Info.Firstname, user.Info.Patronymic, user.Info.Uni, user.Info.Group, user.Info.Grade,
		user.Info.Syllabus, user.Info.City)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *userRepository) GetIDPassword(ctx context.Context, email string) (string, string) {
	var id, pass string
	r.cur.QueryRowContext(ctx, "SELECT id, password FROM users WHERE email=$1", email).Scan(&id, &pass)
	return id, pass
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.cur.ExecContext(ctx, `update info 
	set email=$1, last_name=$2, first_name=$3, patronymic=$4, uni=$5, group=$6, grade=$7, sullabus=$8, city=$9) 
	where user_id=$10`, user.Info.Email,
		user.Info.Lastname, user.Info.Firstname, user.Info.Patronymic, user.Info.Uni, user.Info.Group, user.Info.Grade,
		user.Info.Syllabus, user.Info.City, user.ID)
	return err
}
