package config

import (
	"database/sql"
	"fmt"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func NewMySQLConfig() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	path := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}
	return db, err
}

func NewPosrgreSQLConfig() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")
	path := fmt.Sprintf("postgresql://%s:%s@%s:6543/%s", user, password, host, dbName)
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}
	return db, err
}
