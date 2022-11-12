package models

// NICE_CLASS_LIMIT_SIZE 精选分类模块以及其每个模块显示文章最大条数限制
const NICE_CLASS_LIMIT_SIZE int = 6

// 添加的属性
type addClassVal struct {
	Art        []Article
	ArticleSum int64
	LimitSize  int
}

// SimpleCla 简单提取关键属性
type SimpleCla struct {
	ID   int `gorm:"primary_key"`
	Name string
	addClassVal
}

// Class 映射的类
type Class struct {
	Model
	ParentId    int
	Name        string
	Description string
	Selection   int8
	State       int8
}

// GetClassById 根据id查询分类
func GetClassById(id int) (cla Class) {
	db.Where("id = ?", id).Find(&cla)
	return
}

// GetNiceClass 获取精选分类
func GetNiceClass() (cla []SimpleCla) {
	// 限制一下最大数量为setting.PageSize + 1
	db.Model(&Class{}).Select("id", "name").Where("selection", 1).Limit(NICE_CLASS_LIMIT_SIZE).Find(&cla)
	for i := range cla {
		db.Select("blog_article.id", "blog_article.title", "blog_article.class_id").Joins("left join blog_class c on blog_article.class_id = c.id").Where("c.id = ?", cla[i].ID).Limit(NICE_CLASS_LIMIT_SIZE).Find(&cla[i].Art)
		for j := range cla[i].Art {
			db.Select("id", "name").Where("id = ?", cla[i].Art[j].ClassId).Find(&cla[i].Art[j].Cla)
		}
		db.Model(&Article{}).Joins("left join blog_class c on blog_article.class_id = c.id").Where("c.id = ?", cla[i].ID).Count(&cla[i].ArticleSum)
	}
	return
}

// GetClasses 获取所有分类
func GetClasses() (cla []SimpleCla) {
	db.Model(&Class{}).Select("id", "name").Find(&cla)
	return
}

// GetClaCount 获取分类总数
func GetClaCount() (count int64) {
	db.Model(&Class{}).Count(&count)
	return
}
