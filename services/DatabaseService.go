package services

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func DatabaseService() (*sql.DB, error) {
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		println(err.Error())
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
