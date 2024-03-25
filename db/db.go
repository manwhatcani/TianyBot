package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // 全局数据库连接对象

var sqldb *sqlx.DB

// InitDB 初始化数据库连接
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	CreateTable()
	log.Println("数据库连接成功")
	sqldb = sqlx.NewDb(DB, "sqlite3")
	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// 创建表格
func CreateTable() {
	sqlFile, err := os.ReadFile("table.sql")
	if err != nil {
		log.Fatalf("读取sql文件失败: %v", err)
	}
	sqlStr := string(sqlFile)
	_, err = DB.Exec(sqlStr)
	if err != nil {
		log.Fatalf("执行sql文件失败: %v", err)
	}
	log.Println("读取sql文件成功")
}
