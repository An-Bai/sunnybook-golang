package models

import "sunnybook-golang/pkg/setting"

// 添加的属性
type addArticleVal struct {
	Tags []Tag
	Cla  Class
}

// SimpleArt 简单提取关键属性
type SimpleArt struct {
	Model
	ID      int `gorm:"primary_key"`
	Title   string
	Weather string
}

// Article 映射的类
type Article struct {
	Model
	ClassId        int
	Title          string
	Description    string
	Content        string
	Selection      int8
	State          int8
	WordNum        float64
	ReadOverTime   int
	PreviewContent string
	Author         string
	Weather        string
	addArticleVal
}

// GetArticle 根据id查询具体文章信息
func GetArticle(id int) (art Article) {
	db.Where("id = ?", id).Find(&art)
	// 放入标签、父级包
	ids := GetTagIds(art.ID)
	art.Tags = GetTagsByIds(ids)
	db.Select("id", "name").Where("id", &art.ClassId).Find(&art.Cla)
	return
}

// GetArticleIds 根据标签id查询所有文章id
func GetArticleIds(tid int) (ids []int) {
	db.Table("blog_tag_art").Select("aid").Where("tid = ?", tid).Find(&ids)
	return
}

// GetBaseArticlesPage id数组分页查询该页所有概要文章
func GetBaseArticlesPage(currentPage int, pageSize int, maps interface{}) (articles []Article) {
	offset := (currentPage - 1) * pageSize
	db.Omit("content", "description", "author").Where(maps).Offset(offset).Limit(pageSize).Find(&articles)
	// 放入父级包
	for i := 0; i < len(articles); i++ {
		db.Select("id", "name").Where("id", &articles[i].ClassId).Find(&articles[i].Cla)
	}
	return
}

// GetArticleCount 获取当前查询到的文章总数
func GetArticleCount(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetGreatArticles 查询所有精选概要文章
func GetGreatArticles() (articles []Article) {
	// 限制一下最大数量为setting.PageSize
	num := setting.PageSize
	// 忽略content的字段查询
	db.Omit("content", "preview_content").Where("selection", 1).Limit(num).Find(&articles)
	// 放入标签、父级包
	for i := 0; i < len(articles); i++ {
		ids := GetTagIds(articles[i].ID)
		articles[i].Tags = GetTagsByIds(ids)
		db.Select("id", "name").Where("id", &articles[i].ClassId).Find(&articles[i].Cla)
	}
	return
}

// GetRandomArticles 查询一定数量的随机概要文章
func GetRandomArticles(num int) (articles []Article) {
	// 忽略content的字段查询
	db.Omit("content", "description").Order("Rand()").Limit(num).Find(&articles)
	// 放入父级包
	for i := 0; i < len(articles); i++ {
		db.Select("id", "name").Where("id", &articles[i].ClassId).Find(&articles[i].Cla)
	}
	return
}

// GetRandomArticlesByIds id数组查询所有随机文章
func GetRandomArticlesByIds(ids []int) (articles []Article) {
	// 忽略content的字段查询
	db.Omit("content", "description", "author").Where(ids).Find(&articles)
	// 放入父级包
	for i := 0; i < len(articles); i++ {
		db.Select("id", "name").Where("id", &articles[i].ClassId).Find(&articles[i].Cla)
	}
	return
}

// GetArtCount 获取文章总数
func GetArtCount() (count int64) {
	db.Model(&Article{}).Count(&count)
	return
}

// GetSimpleArts 获取所有文章归档简要信息
func GetSimpleArts() (arts []SimpleArt) {
	db.Order("created_at desc").Model(&Article{}).Select("id", "title", "created_at", "weather").Find(&arts)
	return
}
