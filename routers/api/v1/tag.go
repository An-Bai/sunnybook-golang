package v1

// 大版本 v1.0

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"sunnybook-golang/models"
	error "sunnybook-golang/pkg/e"
)

// GetTagById 获取某个标签
func GetTagById(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	// 先查缓存缓存

	// 没有缓存再查数据库
	data := models.GetTagById(id)
	code := error.SUCCESS
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) {
	// 先查缓存缓存

	// 没有缓存再查数据库
	data := models.GetTags()
	code := error.SUCCESS
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetTagsByIds 获取某篇文章的所有标签
func GetTagsByIds(c *gin.Context) {
	var aid int = -1
	var data []models.Tag
	if arg := c.Param("aid"); arg != "" {
		// 先查缓存缓存
		// 没有缓存再查数据库
		aid = com.StrTo(arg).MustInt()
		tagIds := models.GetTagIds(aid)
		if len(tagIds) > 0 {
			// 添加缓存
			// 继续查文章标签
			data = models.GetTagsByIds(tagIds)
		}
		code := error.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  error.GetMsg(code),
			"data": data,
		})
	}

}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
}

// UpdateTag EditTag 修改文章标签
func UpdateTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}
