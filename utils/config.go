package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TG_BOT_TOKEN string `json:"token"`
	DB_PATH      string `json:"db"`
	PROXY        string `json:"proxy"`
}

func GetConfig() Config {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("无法读取配置文件：%v", err)
	}
	//解析配置文件
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("无法解析配置文件：%v", err)
	}

	return config
}
