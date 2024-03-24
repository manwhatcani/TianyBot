package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/samber/lo"
)

type Link struct {
	id           int
	url          string
	downloadFlag bool
	gid          string
}

func queryHelper(atrbite, table, option string) []Link {
	var sql string
	sql = "SELECT " + atrbite + " " + "FROM " + table
	if !lo.IsEmpty(option) {
		sql = sql + " WHERE " + option
	}

	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []Link
	//TODO 返回多种类型

	for rows.Next() {
		var link Link
		err := rows.Scan(&link)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, link)
		// fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	return results
}

// 获取所有链接
func GetAllLink() []Link {
	allLinks := queryHelper("id,link,downloadFlag", "Links", "")
	// var results []Link
	// for _, data := range allLinks {
	// 	results = append(results, data)
	// }
	return allLinks
}

func GetLinkGidById(id int) string {
	gid := queryHelper("id,gid", "Links", "id = "+strconv.Itoa(id))
	return gid[0].gid
}

func GetLinkById(id int) string {
	gid := queryHelper("id,gid", "Links", "id = "+strconv.Itoa(id))
	return gid[0].gid
}

// 获取所有未下载的链接
func GetNotDownloadLink(oldLinks []string) []string {
	allLinks := queryHelper("link,downloadFlag", "Links", "")
	var results []string
	for _, data := range allLinks {
		if !data.downloadFlag {
			results = append(results, data.url)
		}
	}
	return results
}

func GetUser() {
	// 查询数据
	rows, err := DB.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func CreateTable() {
	// 创建表格
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			age INTEGER
		)`)
	if err != nil {
		log.Println(err)
	}
}

func InsertUser(name string, age int) {
	// 插入数据
	_, err := DB.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, name, age)
	if err != nil {
		log.Println(err)
	}
}

func UpdateLinks() {
	// 查询数据
	rows, err := DB.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Println(err)
		}

	}
}

// 其他 DML 操作类似...
