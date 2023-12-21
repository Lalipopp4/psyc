package user

import (
	"context"
	"psyc/internal/models"
)

func (r *userRepository) Add(ctx context.Context, user *models.User) error {
	_, err := r.cur.ExecContext(ctx, "insert into users (id, last_name, first_name, patronymic, email, age, password, uni, syllabus) values ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Info.Lastname, user.Info.Firstname, user.Info.Patronymic, user.Info.Email, user.Age, user.Password, user.Uni, user.Syllabus)
	return err
}

func (r *userRepository) GetIDPassword(ctx context.Context, email string) (string, string) {
	row := r.cur.QueryRowContext(ctx, "SELECT id, password FROM users WHERE email=$1", email)
	var id, pass string
	row.Scan(&id, &pass)
	return id, pass
}
