package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sunnybook-golang/models"
	"sunnybook-golang/pkg/e"
)

// GetKindsSum 获取某篇文章
func GetKindsSum(c *gin.Context) {
	data := make(map[string]interface{})
	// 缓存没有去查库
	data["tag"] = models.GetTagCount()
	data["article"] = models.GetArtCount()
	data["class"] = models.GetClaCount()
	data["admin"] = "白忆宇"
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
