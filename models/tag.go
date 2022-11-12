package models

// SimpleTag 简单提取关键属性
type SimpleTag struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Tag struct {
	Model
	Name  string
	State int8
}

// GetTagById 根据id查询标签
func GetTagById(id int) (tag Tag) {
	db.Where("id = ?", id).Find(&tag)
	return
}

// GetTags 查询符合条件所有标签
func GetTags() (tags []SimpleTag) {
	db.Model(&Tag{}).Select("id", "name").Find(&tags)
	return
}

// GetTagByName 根据名称查询标签
func GetTagByName(name string) (tags Tag) {
	db.Where("name", name).First(&tags)
	return
}

// GetTagIds 根据文章id查询所有标签id
func GetTagIds(aid int) (tagIds []int) {
	db.Table("blog_tag_art").Select("tid").Where("aid", aid).Find(&tagIds)
	return
}

// GetTagsByIds id数组查询所有标签
func GetTagsByIds(ids []int) (tags []Tag) {
	db.Find(&tags, "id in ?", ids)
	return
}

// GetTagCount 获取标签总数
func GetTagCount() (count int64) {
	db.Model(&Tag{}).Count(&count)
	return
}
