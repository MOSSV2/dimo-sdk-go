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
	ChainType string
	Owner     string
	File      uint64
	Piece     string
	TxHash    string
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

type Stat struct {
	Day           time.Time
	DailyAccounts int64 // new created accounts at this day
	DailyNeedles  int64 // new created needles at this day
	DailyVolumes  int64 // new created volumes at this day
	TotalAccounts int64 // total accounts at this day
	TotalNeedles  int64 // total needles at this day
	TotalVolumes  int64 // total volumes at this day
}

type FileStat struct {
	Day        time.Time
	DailyEdges int64 // new created edges at this day
	DailyFiles int64 // new created files at this day
	DailySize  int64 // new created size at this day
	TotalEdges int64 // total edges at this day
	TotalFiles int64 // total files at this day
	TotalSize  int64 // total size at this day
}
