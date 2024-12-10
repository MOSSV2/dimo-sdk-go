package types

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type GormEdge struct {
	gorm.Model
	Name common.Address
	Type string // "store", "compute"
}

type GormFile struct {
	gorm.Model
	Name  string
	Owner common.Address
}

type GormPiece struct {
	gorm.Model
	Name     string
	Serial   uint64
	File     string
	Index    int // of file
	Owner    common.Address
	Streamer common.Address
}

type GormReplica struct {
	gorm.Model
	Name     string
	Serial   uint64
	Piece    string
	Index    uint8
	StoredOn common.Address
}

type GormEpochProof struct {
	gorm.Model
	Epoch uint64 `gorm:"secondarykey"`
	Store common.Address
}
