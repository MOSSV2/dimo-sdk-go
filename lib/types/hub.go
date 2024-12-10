package types

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name string
}

type Needle struct {
	gorm.Model
	Name   string
	Owner  string
	File   uint64
	Start  uint64
	Size   uint64
	Commit string
}

type Volume struct {
	gorm.Model
	Owner  string
	File   uint64
	Piece  string
	TxHash string
}

type NeedleDisplay struct {
	CreatedAt time.Time
	Name      string
	Owner     string
	File      uint64
	Start     uint64
	Size      uint64
	Piece     string
	TxHash    string
}
