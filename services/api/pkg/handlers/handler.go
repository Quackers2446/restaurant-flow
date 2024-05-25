package handlers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Handler struct {
	DB *sql.DB
}

func New(db *sql.DB) Handler {
	return Handler{db}
}
