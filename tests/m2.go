package tests

import (
	"fmt"
	"regexp"
)

func m2() {
	// 定义正则表达式
	re := regexp.MustCompile(`abv(.*?)xaf`)

	// 要匹配的文本
	text := "afafabvxfsxafghabvxysxaf"

	// 查找所有匹配的子串
	matches := re.FindAllString(text, -1)

	// 输出匹配结果
	for _, match := range matches {
		fmt.Println(match)
	}
}
