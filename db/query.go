package db

import (
	"log"
	"tgbot/model"

	"github.com/samber/lo"
)

// GetLinkByID 根据ID查询用户信息
func GetLinkByID(id int64) (model.Link, error) {
	var link model.Link
	err := sqldb.Get(&link, "SELECT url,downloadFlag,gid FROM links WHERE id=?", id)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

// 获取所有链接
func GetAllLink() ([]model.Link, error) {
	var allLinks []model.Link
	err := sqldb.Select(&allLinks, "SELECT url,downloadFlag,gid FROM links")
	if err != nil {
		log.Printf("获取所有链接失败: %v", err)
		return nil, err
	}
	return allLinks, nil
}

// 获取所有未下载的链接
func GetNotDownloadLink(oldLinks []string) ([]string, error) {
	var allLinks []string
	err := sqldb.Select(&allLinks, "SELECT url FROM links where downloadFlag =true") //这玩意的IN好像不是很好用
	if err != nil {
		return nil, err
	}
	_, reght := lo.Difference[string](allLinks, oldLinks) //取出差集，取右边的
	return reght, nil
}

// 检查链接是否已经下载
func GetDownloadFlagByUrl(url string) (bool, error) {
	var flag bool
	err := sqldb.Get(&flag, "SELECT downloadFlag FROM links WHERE url like ? LIMIT 1", url+"%")
	if err != nil {
		log.Printf("链接获取出错: %v", err)
		return true, err
	}
	return flag, nil
}

// 获取GID
func GetGidByUrl(url string) (string, error) {
	var gid string
	err := sqldb.Get(&gid, "SELECT gid FROM links WHERE url like ? LIMIT 1", url+"%")
	if err != nil {
		log.Printf("链接获取出错: %v", err)
		return "", err
	}
	return gid, nil
}
