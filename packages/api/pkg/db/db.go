package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		user, password, host, port, dbname,
	)

	fmt.Println(connInfo)

	db, err := sqlx.Open("mysql", connInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Applying migrations if any...")
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})

	if err != nil {
		panic(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)

	if err != nil {
		panic(err)
	}

	// Temporary workaround for https://github.com/golang-migrate/migrate/issues/1063
	if err = migrate.Up(); err != nil && err.Error() != "no change" {
		panic(err)
	}
	fmt.Println("Done migrating")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func CloseConnection(db *sqlx.DB) {
	defer db.Close()
}
