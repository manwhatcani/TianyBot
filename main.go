package main

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"tgbot/db"
	"tgbot/model"
	"tgbot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

type GidUrl struct {
	gid string
	url string
}

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

	//初始化数据库
	db.InitDB(config.DB_PATH)
	//关闭数据库连接
	defer db.CloseDB()

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
			msg.Text = isDownloadMsg(update.Message)
		}
		// if lo.IsEmpty(1) {

		// 	// msg.Text = "What can i say,man" + update.Message.Text
		// }
		// log.Printf(update.Message.Text)
		if msg.Text != "" {
			utils.Mylog(msg) //调试用
			continue
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// 检查是否存在下载链接1
func isDownloadMsg(msg *tgbotapi.Message) string {
	//磁链？
	if !lo.IsEmpty(msg.Caption) {
		links := getAllMatchLink(msg.Caption)
		utils.Mylog(links)
		newLinks := getAllNewLink(links)
		utils.Mylog(newLinks)
		results := downloadAll(newLinks)
		gids := lo.Map(results, func(link model.Link, _ int) string {
			return link.Gid
		})
		return "检测到链接，已添加下载: " + strings.Join(gids, ",")
	}
	return "没有检测到链接，跳过。。。"
}

func downloadAll(links []string) []model.Link {

	var gidUrls []model.Link
	for _, url := range links {
		gidUrls = append(gidUrls, model.Link{Gid: utils.SendAria2(url), Url: url, DownloadFlag: true})
	}
	db.CreateLinks(gidUrls)

	return gidUrls

}

// 查询数据库，返回所有新的链接
func getAllNewLink(links []string) []string {
	//TODO
	var newLinks []string
	newLinks, _ = db.GetNotDownloadLink(links)
	return newLinks
}

// 使用正则获取所有的磁链
func getAllMatchLink(caption string) []string {
	re, err := regexp.Compile(`(magnet:\?xt=urn:btih:)[0-9a-fA-F]{40}`)
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
