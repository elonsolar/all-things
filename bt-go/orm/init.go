package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type School struct {
	gorm.Model
	Students []string `gorm:"serializer:json"`
}

func (School) TableName() string {
	return "test_school"
}

var (
	globalDB *gorm.DB
)

func getDB() *gorm.DB {
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

	err = globalDB.AutoMigrate(&School{})
	if err != nil {
		panic(err)
	}
	return globalDB
}
