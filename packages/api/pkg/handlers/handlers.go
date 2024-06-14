package handlers

import (
	"restaurant-flow/pkg/sqlcClient"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB      *sqlx.DB
	Queries *sqlcClient.Queries
}

func New(db *sqlx.DB, queries *sqlcClient.Queries) Handler {
	return Handler{db, queries}
}
