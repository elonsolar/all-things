package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (u User) TableName() string {
	return "test_user"
}
