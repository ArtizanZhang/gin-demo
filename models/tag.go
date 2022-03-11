package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	db.Model(&Tag{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func ExitTagName(tagName string) bool {
	var tag Tag
	db.Model(&Tag{}).Where(map[string]interface{}{"name": tagName}).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false

}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return true
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true

}

// BeforeCreate
// 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
// 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
// 删除：BeforeDelete、AfterDelete
// 查询：AfterFind
func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func GetTag(id int) (tag Tag) {
	db.Where("id = ?", id).First(&tag)
	return
}
