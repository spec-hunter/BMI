package model

import (
	"gorm.io/gorm"
)

type Person struct {
	*gorm.Model
	Name   string
	Weight float64 // 体重
	Height float64 // 身高
	Age    int     // 年龄
	Sex    int     // 性别 (男: 1  女: 0)
	K      float64 // 体脂率
}

func (Person) TableName() string {
	return "person"
}
