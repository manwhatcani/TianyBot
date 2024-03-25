package db

import (
	"log"
	"tgbot/model"
)

// 更新链接的downloadFlag信息
func UpdateLink(downloadFlag bool, url string) error {
	_, err := sqldb.Exec("UPDATE links SET downloadFlag=? WHERE url=?", downloadFlag, url)
	// _, err := sqlx.NamedExec(sqldb, "UPDATE links SET username=:username WHERE id=:id", link)
	if err != nil {
		log.Printf("更新下载url(%s)状态为:%b 失败:%v", url, err, downloadFlag)
		return err
	}
	log.Println("更新下载状态成功")
	return nil
}

// func CreateLinks(link []*model.Link) ([]int, error) {
// 	var result []int
// 	data, err := sqldb.NamedExec("INSERT INTO links (url, downloadFlag,gid) VALUES (:url, :downloadFlag, :gid)", link)
// 	utils.Mylog(data)
// 	if err != nil {
// 		log.Printf("插入数据库失败:%v", err)
// 		return nil, err
// 	}
// 	return result, nil
// }

// 插入数据
// TODO 返回出错的gid，便于后续处理
func CreateLinks(links []model.Link) error {
	for _, link := range links {
		result, err := sqldb.NamedExec("INSERT INTO links (url, downloadFlag, gid) VALUES (:url, :downloadFlag, :gid)", link)
		if err != nil {
			log.Printf("数据库插入 %s 失败: %v", link.Url, err)
			continue
		}
		log.Println("结果:", result)
	}
	return nil
}
