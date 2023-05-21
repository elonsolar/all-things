package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// gorm 中的where
// where 子句从语法层面讲， 是sql 中最灵活，最复杂的

//1. gorm where 的基本使用
// 1. 字符串 灵活，  转为 clause.Expr{SQL: s, Vars: args} ,clause.Expression 的一种
// 2. map ,结构体    ，   只能表达  eq ,in  BuildCondition
// 3. clause.Expression , gorm sql底层表示

// Where  所有的条件 以 []Expression 表示， 【Expression 之间默认是 “AND” 连接】
// Expression, 有Like, Eq, Expr , AND ,OR ...

//type Where struct {
//Exprs []Expression
// }

// type AndConditions struct {
//	Exprs []Expression
// }

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (u User) TableName() string {
	return "test_user"
}
func TestSimple(t *testing.T) {

	users := []User{}

	//SELECT * FROM `test_user` WHERE name like '%xm%'  AND age in (12,13) AND `test_user`.`deleted_at` IS NULL
	// err := getDB().Where("name like ? ", "%xm%").Where("age in ?", []int{12, 13}).Find(&users).Error
	err := getDB().Where("name like ? and age in ? ", "%xm%", []int{12, 13}).Find(&users).Error

	if err != nil {
		panic(err)
	}
	log.Println(users)
}

func TestStructMap(t *testing.T) {

	users := []User{}

	params := map[string]interface{}{
		"name": "xm",
		"age":  []int{12, 13},
	}
	params2 := &User{Age: 25}

	//SELECT * FROM `test_user` WHERE `age` IN (12,13) AND `name` = 'xm' AND `test_user`.`deleted_at` IS NULL
	err := getDB().Where(params).Where(params2).Find(&users).Error

	if err != nil {
		panic(err)
	}
	log.Println(users)
}

func TestClause(t *testing.T) {

	users := []User{}

	condEQ := clause.Eq{Column: "age", Value: 14}

	condLike := clause.Like{Column: "name", Value: "%xm%"}
	clause.And()

	// SELECT * FROM `test_user` WHERE `age` = 14 AND `name` LIKE '%xm%' AND `test_user`.`deleted_at` IS NULL

	// err := getDB().Where(condEQ).Where(condLike).Find(&users).Error // 这种场景 这两种是一样的
	err := getDB().Clauses(condEQ, condLike).Find(&users).Error // 这种场景 这两种是一样的
	/**
	  func (db *DB) Clauses(conds ...clause.Expression) (tx *DB) {
	  	tx = db.getInstance()
	  	var whereConds []interface{}

	  	for _, cond := range conds {
	  		if c, ok := cond.(clause.Interface); ok {
	  			tx.Statement.AddClause(c)
	  		} else if optimizer, ok := cond.(StatementModifier); ok {
	  			optimizer.ModifyStatement(tx.Statement)
	  		} else {
	  			whereConds = append(whereConds, cond)
	  		}
	  	}

	  	if len(whereConds) > 0 {
	  		tx.Statement.AddClause(clause.Where{Exprs: tx.Statement.BuildCondition(whereConds[0], whereConds[1:]...)})
	  	}
	  	return


	*/

	if err != nil {
		panic(err)
	}
	fmt.Println(users)

}
