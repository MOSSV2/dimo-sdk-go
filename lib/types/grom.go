package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type Strs []string

func (s *Strs) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan")
	}
	return json.Unmarshal(bytesValue, s)
}

func (s Strs) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type GormEdge struct {
	gorm.Model
	Name      common.Address
	Type      string // "store", "compute"
	ChainType string
	PublicKey []byte
	ExposeURL string
	CPU       string
	Memory    string
	OnChain   bool
	Last      time.Time
	Revenue   string
}

type GormChain struct {
	gorm.Model
	Name   string
	Height uint64
}

type GormFile struct {
	gorm.Model
	RSN       int
	RSK       int
	Name      string
	Hash      string
	Size      int64
	Owner     common.Address
	Pieces    Strs
	ChainType string
}

type GormPiece struct {
	gorm.Model
	RSN       int
	RSK       int
	Name      string
	Serial    uint64
	Size      int64
	Start     uint64
	Expire    uint64
	Price     string
	Owner     common.Address
	Streamer  common.Address
	TxHash    string
	ChainType string
}

type GormReplica struct {
	gorm.Model
	Fake      bool
	Name      string
	Serial    uint64
	Size      int64
	Piece     string
	Index     uint8
	StoredOn  common.Address
	Ordinal   uint64
	TxHash    string
	ChainType string
}

type GormEpochProof struct {
	gorm.Model
	Fake      bool
	Epoch     uint64
	Store     common.Address
	TxHash    string
	Hash      []byte
	Sum       []byte
	H         []byte
	Value     []byte
	ChainType string
}

type GormReplicaWitness struct {
	gorm.Model
	Serial    uint64
	Proof     []byte
	ChainType string
}

type GormStat struct {
	gorm.Model
	Name      string
	Count     int64
	Size      int64
	ChainType string
}
