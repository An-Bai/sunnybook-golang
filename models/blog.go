package models

import "time"

type Self struct {
	Name         string
	Author       string
	About        string
	AboutBlog    string
	AboutAuthor  string
	AboutSetting string
	BornAt       time.Time
}

// GetArticleAbout 获取
func GetArticleAbout() (blog Self) {
	db.First(&blog)
	return
}
