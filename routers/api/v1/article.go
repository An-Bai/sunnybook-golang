package v1

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
	"strconv"
	"sunnybook-golang/models"
	error "sunnybook-golang/pkg/e"
	"sunnybook-golang/pkg/setting"
	"sunnybook-golang/pkg/util"
	"time"
)

// GetArticleById 获取某篇文章
func GetArticleById(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	// 直接从缓存中获取

	// 缓存没有去查库
	data := models.GetArticle(id)
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetGreatArticles 显示精选分类文章
func GetGreatArticles(c *gin.Context) {
	// 直接从缓存中获取

	// 缓存没有去查库
	data := models.GetGreatArticles()
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetRandomArticles 获取随机文章
func GetRandomArticles(c *gin.Context) {
	var code int
	// 直接从random.ini中获取随机文章的ids
	cfg, err := ini.Load("conf/random.ini")
	if err != nil {
		// 打印日志输出信息
		log.Printf("Fail to parse 'conf/random.ini': %v", err)
		code = error.ERROR
	}
	// 获取当天时间
	dateStr := time.Now().Format("06-01-02")
	// 获取记录的随机文章刷新时间
	s := cfg.Section("").Key("LAST_TIME").String()
	// 比较，不超过当天就取，否则要刷新ids
	if dateStr == s {
		// go-ini 的方法，从文件中获取指定字符串格式的[]int
		ids := cfg.Section("").Key("RANDOM_ARTICLE_ID").Ints(",")
		// 如果拿到了数据就返回请求响应（一般来说都会有数据，除非在读ids []int的时候还没写出来，那么其实可以写一下循环模拟自旋一定毫秒判断一下）
		if ids != nil && len(ids) > 0 {
			data := models.GetRandomArticlesByIds(ids)
			code = error.SUCCESS
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  error.GetMsg(code),
				"data": data,
			})
			return
		}
	}
	// 缓存没有去查库
	// 获取文章总数
	var maps map[string]interface{}
	count := models.GetArticleCount(maps)
	num := setting.PageSize
	// 小于默认10就替换默认值
	if count < int64(num) {
		num = int(count)
	}
	// 获取一定数量的随机文章
	data := models.GetRandomArticles(num)
	//获取ids
	ids := make([]int, num)
	for i, article := range data {
		ids[i] = article.ID
	}
	// 开启写文件的协程，写入当天时间和文章的id
	go writeRecord(dateStr, ids)
	// 直接返回查到的文章
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetBaseArticlesPage 获取分页文章
func GetBaseArticlesPage(c *gin.Context) {
	var page util.Page
	//query := c.Query("currentPage")
	//println(query)
	if err := c.Bind(&page); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"code": error.INVALID_PARAMS,
			"msg":  error.GetMsg(error.INVALID_PARAMS),
			"data": "",
		})
		return
	}
	maps := make(map[string]interface{})
	// 类和标签分页筛选
	if classId := c.Query("classId"); classId != "" {
		maps["class_id"] = classId
	} else if tagId := c.Query("tagId"); tagId != "" {
		tid, _ := strconv.Atoi(tagId)
		maps["id"] = models.GetArticleIds(tid)
	}
	// 获取分页总数
	count := models.GetArticleCount(maps)
	// 获取分页数据
	page.PageData = models.GetBaseArticlesPage(page.CurrentPage, setting.PageSize, maps)
	page.Count = int(count)
	page.PageSize = setting.PageSize
	c.JSON(http.StatusOK, gin.H{
		"code": error.SUCCESS,
		"msg":  error.GetMsg(error.SUCCESS),
		"data": page,
	})
}

// GetSimpleArts 获取所有文章并进行归档
func GetSimpleArts(c *gin.Context) {
	// 直接从缓存中获取

	// 缓存没有去查库
	arts := models.GetSimpleArts()

	signYear := arts[0].CreatedAt.Year()
	signMonth := arts[0].CreatedAt.Month()
	var data [][][]models.SimpleArt
	i := 0
	j := 0
	k := 0
	data = make([][][]models.SimpleArt, 1)
	data[i] = make([][]models.SimpleArt, 1)
	for _, art := range arts {
		if art.CreatedAt.Year() != signYear {
			signYear = art.CreatedAt.Year()
			i++
			j = 0
			k = 0
			data = append(data, make([][]models.SimpleArt, 1))
		} else if art.CreatedAt.Month() != signMonth {
			signMonth = art.CreatedAt.Month()
			j++
			k = 0
			data[i] = append(data[i], make([]models.SimpleArt, 1))
		}
		if k == 0 {
			data[i][j] = make([]models.SimpleArt, 1)
			data[i][j][0] = art
		} else {
			data[i][j] = append(data[i][j], art)
		}
		k++
	}
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// GetArticleAbout 获取文章关于内容
func GetArticleAbout(c *gin.Context) {
	// 直接从缓存中获取

	// 缓存没有去查库
	data := models.GetArticleAbout()
	code := error.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}

// 功能方法区
// 将随机文章的当天信息写入文件（一天执行一次左右）
func writeRecord(date string, ids []int) {
	// 先清空，权限：清空
	file, err := os.OpenFile("conf/random.ini", os.O_RDWR|os.O_TRUNC, 0775)
	if err != nil {
		log.Printf("打开文件variation.ini失败,err: + %v", err)
	}
	// 再写入，权限：读写和追加
	file, _ = os.OpenFile("conf/random.ini", os.O_RDWR|os.O_APPEND, 0775)
	defer file.Close()
	//ids := []int{1, 2, 3, 4, 5}
	var buffer bytes.Buffer
	for i, id := range ids {
		fmt.Printf("%d,", id)
		//println(strconv.Itoa(id))
		buffer.WriteString(strconv.Itoa(id))
		if i < len(ids)-1 {
			buffer.WriteString(",")
		}
	}
	//println(buffer.String())
	file.WriteString(fmt.Sprintf("LAST_TIME = %s\n", date))
	file.WriteString(fmt.Sprintf("RANDOM_ARTICLE_ID = %s\n", buffer.String()))
}
