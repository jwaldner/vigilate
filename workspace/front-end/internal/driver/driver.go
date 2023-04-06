package driver

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn" // need this and next two for pgx
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection information
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 25
const maxIdleDbConn = 25
const maxDbLifetime = 5 * time.Minute

var counts int64

// code from micro services
func openDB(dsn string) (*DB, error) {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)
	dbConn.SQL = db

	return dbConn, nil
}

func ConnectToDB() *DB {

	dsn := os.Getenv("DSN")

	log.Println("my dsn",dsn)


	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
			counts++
		} else {
			log.Println("connected to postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

