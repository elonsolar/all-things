package main

import (
	"fmt"
	"generic/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var (
	globalDB *gorm.DB
)

type dao[T any, K any] struct {
}

func (d dao[T, K]) create(t *T) {
	globalDB.Create(t)
}

func (d dao[T, K]) batchCreate(t *[]T) {
	globalDB.Create(t)
}

func (dao[T, K]) updateSelectiveByParams(t *T, params clause.Expression) (int64, error) {

	result := globalDB.Where(params).Updates(t)

	return result.RowsAffected, result.Error
}

func (d dao[T, K]) updateSelectiveById(t *T) (int64, error) {
	result := globalDB.Model(t).Updates(t)

	return result.RowsAffected, result.Error
}

func (d dao[T, K]) deleteById(id K) (int64, error) {
	result := globalDB.Delete(new(T), id)
	return result.RowsAffected, result.Error
}

type userDao[u model.User, id uint] struct {
	dao[u, id]
}

func main() {

	initMysql()
	// g := gen.NewGenerator(gen.Config{
	// 	OutPath: "/Users/wendy/study/repo/all-things/bt-go/generic/query",
	// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	// })

	// g.UseDB(globalDB)

	// g.ApplyBasic(model.User{})
	// g.ApplyInterface(func(Querier) {}, model.User{})

	// g.Execute()

	// query.SetDefault(globalDB)
	// user, err := query.User.Where(query.User.Name.Eq("xm")).First()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)

	// ua := user{gorm.Model{
	// 	ID: 3,
	// }, "mht", 18}

	// ub := user{gorm.Model{
	// 	ID: 4,
	// }, "xm", 22}
	userdao := userDao[model.User, uint]{}

	// userdao.batchCreate(&[]user{ua, ub})

	// ClauseTest()

	updateU := &model.User{Age: 40}

	// params := clause.Like{Column: "name", Value: "%ibt%"}
	params := clause.Or(
		clause.Like{Column: "name", Value: "%hah%"},
		clause.Like{Column: "id", Value: "%hah%"},
		clause.And(
			clause.Eq{Column: "age", Value: 12},
			clause.Eq{Column: "name", Value: "gg"},
		),
	)

	updateCount, err := userdao.updateSelectiveByParams(updateU, params)
	// userdao.updates(&u)
	// updateCount, err := userdao.deleteById(1)
	if err != nil {
		panic(err)
	}
	log.Printf("更新数:%d", updateCount)
}

func initMysql() {
	mysqlConfig := mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         191,                                                                            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                          // 根据版本自动配置
	}
	var err error
	globalDB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	globalDB.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, _ := globalDB.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(2)

	err = globalDB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}

func retErr() error {

	var err *myError
	fmt.Println("retError err ==nil", err == nil)
	return err
}

type myError struct {
}

func (myError) Error() string {
	return "ss"
}

func mainV() {

	var err = retErr()

	fmt.Println("main err==nil", err == nil) //false
}

func ClauseTest() {
	stmt := &gorm.Statement{DB: globalDB, Table: "test_user", Clauses: map[string]clause.Clause{}}
	clause.Expr{SQL: "create table ? (? ?, ? ?)", Vars: []interface{}{clause.Table{Name: "ggg"}, clause.Column{Name: "id"}, clause.Expr{SQL: "int"}, clause.Column{Name: "name"}, clause.Expr{SQL: "text"}}}.Build(stmt)

	fmt.Println(stmt.SQL.String())

}
