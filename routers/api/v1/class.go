package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sunnybook-golang/models"
	"sunnybook-golang/pkg/e"
)

// GetNiceClass 获取精选分类
func GetNiceClass(c *gin.Context) {
	// 直接从缓存中获取

	// 缓存没有去查库
	data := models.GetNiceClass()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetClasses 获取所有分类
func GetClasses(c *gin.Context) {
	// 直接从缓存中获取

	// 缓存没有去查库
	data := models.GetClasses()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetClassById 获取id对应分类
func GetClassById(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	// 直接从缓存中获取
	// 缓存没有去查库
	data := models.GetClassById(id)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
