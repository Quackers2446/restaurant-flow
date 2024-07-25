package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

const (
	host     = "localhost"
	port     = 3307
	user     = "user"
	password = "password"
	dbname   = "restaurantFlowAuth"
)

func Connect() *sqlx.DB {
	connInfo := fmt.Sprintf(
		// "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true&parseTime=true",
		user, password, host, port, dbname,
	)

	fmt.Println(connInfo)

	db, err := sql.Open("mysql", connInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Applying migrations if any...")
	driver, err := mysql.WithInstance(db, &mysql.Config{})

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

	// handle err
	loggerAdapter := zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}))

	return sqlx.NewDb(sqldblogger.OpenDriver(connInfo, db.Driver(), loggerAdapter), "mysql")
}

func CloseConnection(db *sqlx.DB) {
	defer db.Close()
}
