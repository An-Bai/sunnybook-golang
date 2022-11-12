package util

import "sunnybook-golang/models"

type Page struct {
	PageSize    int `json:"pageSize"`
	CurrentPage int `form:"currentPage" params:"currentPage"  json:"currentPage" uri:"currentPage" xml:"currentPage" binding:"required"`
	Count       int `json:"count"`
	PageData    []models.Article
}
