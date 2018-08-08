package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
}

var DefaultDB = &DB{
	"root", "root", "test", "localhost", "3306",
}

func (db *DB) Connection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		db.User, db.Password, db.Host, db.Port, db.DBName))
}
