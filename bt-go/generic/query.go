package main

import (
	"gorm.io/gen"
)

type Querier interface {

	// select * from @@table where name =@name{{if role !=""}} and role =@role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}
