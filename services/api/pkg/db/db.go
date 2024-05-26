package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "user"
	password = "password"
	dbname   = "restaurantFlow"
)

func Connect() *sqlx.DB {
	connInfo := fmt.Sprintf(
		// "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname,
	)

	fmt.Println(connInfo)

	db, err := sqlx.Open("mysql", connInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func CloseConnection(db *sqlx.DB) {
	defer db.Close()
}
