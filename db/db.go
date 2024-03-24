package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	// 打开数据库连接
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Println(err)
	}
	DB = db
}
