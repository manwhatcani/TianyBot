package main

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"tgbot/db"
	"tgbot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

func command(com string) string {
	// Extract the command from the Message.
	msg := ""
	switch com {
	case "help":
		msg = "I understand /sayhi and /status."
	case "sayhi":
		msg = "Hi :)"
	case "status":
		msg = "I'm ok."
	default:
		msg = "I don't know that command"
	}
	return msg
}

func main() {
	config := utils.GetConfig()

	var proxyStr string
	if config.PROXY != "" {
		proxyStr = config.PROXY // 替换为您的代理地址
	}

	// 设置代理
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{
		Transport: transport,
	}

	// bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	TG_BOT_TOKEN := config.TG_BOT_TOKEN
	// bot, err := tgbotapi.NewBotAPI(TG_BOT_TOKEN)
	bot, err := tgbotapi.NewBotAPIWithClient(TG_BOT_TOKEN, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() { // ignore any non-command Messages
			msg.Text = command(update.Message.Command())
		} else {
			isDownloadMsg(update.Message)
		}
		// if lo.IsEmpty(1) {

		// 	// msg.Text = "What can i say,man" + update.Message.Text
		// }
		// log.Printf(update.Message.Text)
		if msg.Text != "" {
			utils.Mylog(msg)
			continue
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// 检查是否存在下载链接1
func isDownloadMsg(msg *tgbotapi.Message) {
	//磁链？
	if !lo.IsEmpty(msg.Caption) {
		links := getAllMatchLink(msg.Caption)
		utils.Mylog(links)
		newLinks := getAllNewLink(links)
		utils.Mylog(newLinks)
		// downloadAll(newLinks);
	}
}

// 查询数据库，返回所有新的链接
func getAllNewLink(links []string) []string {
	db.Init()
	db.CreateTable()
	db.GetUser()
	//TODO
	var newLinks []string
	db.GetNotDownloadLink(links)
	return newLinks
}

// 使用正则获取所有的磁链
func getAllMatchLink(caption string) []string {
	re, err := regexp.Compile("(magnet:.*?)\n")
	if err != nil {
		log.Panic(err)
	}
	// 使用FindAllString方法查找所有匹配的字符串
	matches := re.FindAllStringSubmatch(caption, -1)
	// log.Println("全部内容:", matches) // 输出: All matches: [123]
	var links []string
	if len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 0 {
				links = append(links, strings.Replace(match[0], `\n`, "", -1)) //替换带\n的链接
			}
		}
	}
	return links
}
