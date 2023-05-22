package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func createSchool(school *School) error {

	return getDB().Create(school).Error
}

func selectSchoolList() ([]School, error) {

	schoolList := make([]School, 0, 10)
	err := getDB().Where("id", []interface{}{1, 2}).Find(&schoolList).Error
	return schoolList, err
}

func TestCreateSchool(t *testing.T) {
	school := School{gorm.Model{ID: 1}, []string{"xm", "xh"}}
	createSchool(&school)
}

func TestFind(t *testing.T) {

	schoolList, err := selectSchoolList()
	if err != nil {
		panic(err)
	}

	fmt.Println(schoolList)

}
