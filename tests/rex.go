package tests

import (
	"log"
	"regexp"
)

func rex() {
	// 定义要匹配的字符串
	text := "2024-03-22\nFC2-PPV-4343600\n※緊急【期間限定3/31まで】18歳なりたて学校卒業直後の純白スレンダー天使♥人生初のハメ撮りで生挿入、 中出し2連発！美しく括れた神ウエストでパイパン美マンに大量膣内射精されて放心状態！\n\nmagnet:?xt=urn:btih:abbb693cffacd94f4592fbb417e41338fcaca507\n\n#t2024_03_22 #ハメ撮りマスター「D」 #sangou #スレンダー #美女 #可愛い #18歳 #卒業したて2024-03-22\nFC2-PPV-4343600\n※緊急【期間限定3/31まで】18歳なりたて学校卒業直後の純白スレンダー天使♥人生初のハメ撮りで生挿入、 中出し2連発！美しく括れた神ウエストでパイパン美マンに大量膣内射精されて放心状態！\n\nmagnet:?xt=urn:btih:abbb693cffacd94f4592fbb417e41338fcaca507\n\n#t2024_03_22 #ハメ撮りマスター「D」 #sangou #スレンダー #美女 #可愛い #18歳 #卒業したて"

	log.Println("开始匹配：", text)
	// 定义正则表达式
	// 这个正则匹配一个或多个数字
	// re, err := regexp.Compile(`\d+`)
	re, err := regexp.Compile("(magnet:.*?)\n")
	if err != nil {
		log.Panic(err)
	}

	// 使用FindString方法查找匹配的字符串
	match := re.FindString(text)
	log.Println("单个匹配:", match) // 输出: Match found: 123

	// 使用FindAllString方法查找所有匹配的字符串
	matches := re.FindAllStringSubmatch(text, -1)
	log.Println("全部内容:", matches) // 输出: All matches: [123]

	// // 使用MatchString方法检查字符串是否匹配正则表达式
	// matched, err := regexp.MatchString(`\w+`, text)
	// if err != nil {
	// 	fmt.Println("Regex compilation error:", err)
	// } else {
	// 	fmt.Println("Is text matched by regex?", matched) // 输出: Is text matched by regex? true
	// }

	// // 使用Submatch方法查找匹配的子表达式
	// reSub := regexp.MustCompile(`(\w+)\s(\w+)`)
	// matchesSub := reSub.FindStringSubmatch(text)
	// fmt.Println("Submatches:", matchesSub) // 输出: Submatches: [Hello World Hello World]
}
