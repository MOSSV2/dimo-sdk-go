package types

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name string
}

type Needle struct {
	gorm.Model
	Name  string
	Owner string
	File  uint32
	Start uint64
	Size  uint64
}
