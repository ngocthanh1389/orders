package dataaccess

import "github.com/jmoiron/sqlx"

type OrderDAL struct {
	db *sqlx.DB
}

func NewOrderDAL(db *sqlx.DB) *OrderDAL {
	return &OrderDAL{
		db: db,
	}
}
