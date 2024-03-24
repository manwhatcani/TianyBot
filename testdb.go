package main

import (
	"tgbot/db"
)

func testdb() {
	db.Init()
	db.CreateTable()
	db.InsertUser("张三", 18)
	db.GetUser()
	db.DB.Close()
}
