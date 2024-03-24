package utils

import (
	"encoding/json"
	"log"
)

func Mylog(person any) {
	// 将结构体编码为 JSON 格式的字节流
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// 将字节流转换为字符串
	jsonString := string(jsonData)
	log.Println("tgbot:", jsonString) // 输出: {"name":"Alice","age":30}
}
