package main

import (
	"testing"

	"gorm.io/gorm"
)

// update

// 1. model 也是条件?  是否可以不写？
// 1.1 除了 where ,clause  等子句， model 里的主键也会被拿来作为条件

// user.id=2
// updateUser.name=""
//UPDATE `test_user` SET `id`=2,`name`='xm' WHERE age=11 AND `test_user`.`deleted_at` IS NULL AND `id` = 2
// db.model(&user).where("age=?",11).updates(updateUser)
//
// 1.2  如果 updates(结构体) ,就可以省略 不写，但是 要有where 条件，或者开启全局更新？

// 2. 更新 多个字段
// 2.1 更新单个字段
// db.model(&user).update("age",11)
// 2.2 更新多个字段，map, struct
// db.model(&user).updates(map[string]interface{}{})
// db.model(&user).updates(user)

// 3. 全局更新
// db.updateColumn

// 4. 选择/忽略 更新
// omit,select

type People struct {
	gorm.Model
}

func (People) TableName() string {
	return "people"
}

func TestUpdate(t *testing.T) {

	usr := User{Name: "xm"}
	usr.ID = 2
	getDB().Model(usr).Update("age", 11)
}
