package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"sunnybook-golang/pkg/setting"
)

// GetPage 从请求中获取当前页面
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
