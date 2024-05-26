package handlers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) Handler {
	return Handler{db}
}
