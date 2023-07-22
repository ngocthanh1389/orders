package dataaccess

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserDAL struct {
	db *sqlx.DB
}

func NewUserDAL(db *sqlx.DB) *UserDAL {
	return &UserDAL{
		db: db,
	}
}

func (u *UserDAL) ValidateUser(name, pass string) (bool, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("count(*)").
		From("users").Where(squirrel.Eq{"name": name, "pass": pass}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("build query error: %w", err)
	}

	t := 0
	if err := u.db.Get(&t, q, p...); err != nil {
		return false, fmt.Errorf("run query error: %w", err)
	}
	return t > 0, nil
}
