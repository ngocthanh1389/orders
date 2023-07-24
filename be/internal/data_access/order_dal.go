package dataaccess

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type OrderDAL struct {
	db *sqlx.DB
}

func NewOrderDAL(db *sqlx.DB) *OrderDAL {
	return &OrderDAL{
		db: db,
	}
}

func (o *OrderDAL) CheckHasOrderToday(name string) (bool, error) {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("max(time)").
		From("orders").
		Where(squirrel.Eq{"name": name}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("build query error: %w", err)
	}

	var t *time.Time
	if err := o.db.Get(&t, q, p...); err != nil {
		return false, fmt.Errorf("run query error: %w", err)
	}
	if t == nil {
		return false, nil
	}
	year, month, day := t.Local().Date()
	curYear, curMonth, curDay := time.Now().Local().Date()
	if year == curYear && month == curMonth && day == curDay {
		return true, nil
	}
	return false, nil
}

func (o *OrderDAL) Insert(name string, t time.Time) error {
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("orders").
		Columns("name", "time").
		Values(name, t).
		ToSql()
	if err != nil {
		return fmt.Errorf("build query error: %w", err)
	}

	if _, err := o.db.Exec(q, p...); err != nil {
		return fmt.Errorf("run query error: %w", err)
	}
	return nil
}
