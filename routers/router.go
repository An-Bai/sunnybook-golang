package routers

import (
	"github.com/gin-gonic/gin"
	"sunnybook-golang/pkg/setting"
	"sunnybook-golang/pkg/util"
	v1 "sunnybook-golang/routers/api/v1"
)

func InitRouter() *gin.Engine {
	//router := gin.Default()
	// 手动调用gin.Default()里面的默认实现（一样的效果），顺便添加上运行模式配置
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(util.Cors())
	// 配置运行模式
	gin.SetMode(setting.RunMode)

	// 测试路由
	//r.GET("/hello", func(c *gin.Context) {
	//	tags := models.GetTags()
	//	fmt.Println(tags)
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": 100,
	//		"data": tags,
	//	})
	//})

	apiv1 := r.Group("/api/v1")
	{
		// 标签
		{
			//获取标签列表
			apiv1.GET("/tags", v1.GetTags)
			//获取某个标签
			apiv1.GET("/tag/:id", v1.GetTagById)

			//新建标签
			apiv1.POST("/tag", v1.AddTag)
			//更新指定标签
			apiv1.PUT("/tag/:id", v1.UpdateTag)
			//删除指定标签
			apiv1.DELETE("/tag/:id", v1.DeleteTag)
		}
		// 文章
		{
			// 获取某篇文章
			apiv1.GET("/article/:id", v1.GetArticleById)
			// 获取精选概要文章
			apiv1.GET("/article/great", v1.GetGreatArticles)
			// 获取每日随机概要文章
			apiv1.GET("/article/random", v1.GetRandomArticles)
			// 获取分页文章
			apiv1.GET("/article/page", v1.GetBaseArticlesPage)
			// 获取归档文章
			apiv1.GET("/article/pigeonhole", v1.GetSimpleArts)
		}
		// 分类
		{
			// 获取精选概要分类
			apiv1.GET("/class/nice", v1.GetNiceClass)
			// 获取所有分类
			apiv1.GET("/class", v1.GetClasses)
			apiv1.GET("/class/:id", v1.GetClassById)
		}
		// 其它接口
		{
			// 获取精选概要分类
			apiv1.GET("/blog/kindSum", v1.GetKindsSum)
			// 获取关于文章
			apiv1.GET("/blog/about", v1.GetArticleAbout)
		}
	}
	return r
}
