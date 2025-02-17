package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag xxx
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// EditTag xxx
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

// DeleteTag xxx
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

//BeforeCreate xxx
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_on", time.Now().Unix())
	fmt.Println("BeforeCreate")
	return nil
}

// BeforeUpdate xxx
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on", time.Now().Unix())
	fmt.Println("BeforeUpdate")
	return nil
}
