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

type OrderEntity struct {
	ID   int64     `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
	Time time.Time `db:"time" json:"time"`
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

func (o *OrderDAL) GetTodayOrder() ([]OrderEntity, error) {
	curYear, curMonth, curDay := time.Now().Local().Date()
	start := time.Date(curYear, curMonth, curDay, 0, 0, 0, 0, time.Local)
	end := time.Date(curYear, curMonth, curDay, 23, 59, 59, 0, time.Local)
	q, p, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("*").
		From("orders").
		Where(squirrel.And{squirrel.Gt{"time": start}, squirrel.Lt{"time": end}}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query error: %w", err)
	}

	var res []OrderEntity
	if err := o.db.Select(&res, q, p...); err != nil {
		return nil, fmt.Errorf("run query error: %w", err)
	}
	return res, nil
}
