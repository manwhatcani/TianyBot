package model

type Link struct {
	Url          string `db:"url"`
	DownloadFlag bool   `db:"downloadFlag"`
	Gid          string `db:"gid"`
}

// func (link Link) Value() (driver.Value, error) {
// 	return []interface{}{link.Url, link.DownloadFlag, link.Gid}, nil
// }

type BaseLink struct {
	Id           int    `db:"id"`
	Url          string `db:"url"`
	DownloadFlag bool   `db:"downloadFlag"`
	Gid          string `db:"gid"`
}

type LinkUrlDownloadFlag struct {
	Url          string `db:"url"`
	DownloadFlag bool   `db:"downloadFlag"`
}

type LinkUrlGid struct {
	Url string `db:"url"`
	Gid string `db:"gid"`
}
