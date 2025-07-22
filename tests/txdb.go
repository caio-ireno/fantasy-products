package tests

import (
	"app/internal/application"
	"database/sql"
	"os"
	"sync"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var registerOnce sync.Once

func RegisterDatabase() {
	registerOnce.Do(func() {

		err := godotenv.Load("../../.env.test")
		if err != nil {
			panic("error loading .env file")
		}

		cfg := &application.ConfigApplicationDefault{
			Db: &mysql.Config{
				User:   os.Getenv("DB_USER"),
				Passwd: os.Getenv("DB_PASSWORD"),
				Net:    "tcp",
				Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
				DBName: os.Getenv("DB_NAME"),
			},
		}

		dsn := cfg.Db.FormatDSN()
		txdb.Register("txdb", "mysql", dsn)
	})
}

func InitDatabase() error {
	conn, err := sql.Open("txdb", "identifier")

	db = conn

	return err
}

func GetDB() *sql.DB {
	return db
}

func init() {
	RegisterDatabase()
}
