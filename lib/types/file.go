package types

import (
	"context"
	"io"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type Policy struct {
	N, K int8
}

type FileCore struct {
	OnChain  bool
	Owner    common.Address
	Name     string
	Hash     string
	Policy   Policy
	Size     int64
	Creation time.Time
	Pieces   []string
}

func (frc *FileCore) Serialize() ([]byte, error) {
	return cbor.Marshal(frc)
}

func (frc *FileCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, frc)
}

type FileReceipt struct {
	FileCore
}

func (fr *FileReceipt) Serialize() ([]byte, error) {
	return cbor.Marshal(fr)
}

func (fr *FileReceipt) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, fr)
}

type FileCoreWithSize struct {
	FileCore
	OnlyPiece bool
	Sizes     []int64
}

func (fcws *FileCoreWithSize) Serialize() ([]byte, error) {
	return cbor.Marshal(fcws)
}

func (fcws *FileCoreWithSize) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, fcws)
}

type PieceCore struct {
	Policy Policy
	Name   string
	Serial int64
	Size   int64 // raw size
	Price  *big.Int
}

func (pc *PieceCore) Serialize() ([]byte, error) {
	return cbor.Marshal(pc)
}

func (pc *PieceCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, pc)
}

type PieceReceipt struct {
	PieceCore
	Replicas []string
	StoredOn []common.Address
}

func (cr *PieceReceipt) Serialize() ([]byte, error) {
	return cbor.Marshal(cr)
}

func (cr *PieceReceipt) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, cr)
}

type ReplicaCore struct {
	Name     string // encoded
	Size     int64  // stored size
	Piece    string // belongs to which piece
	Index    uint8
	StoredOn common.Address
}

func (rc *ReplicaCore) Serialize() ([]byte, error) {
	return cbor.Marshal(rc)
}

func (rc *ReplicaCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, rc)
}

type ReplicaInfo struct {
	File    uint64
	Piece   uint64
	Index   int
	Replica string
	Sign    []byte
	Store   common.Address
}

type IFile interface {
	Put(context.Context, common.Address, io.Reader) (FileCoreWithSize, error)
	Get(context.Context, string, io.Writer) (FileReceipt, error)
	GetPiece(context.Context, string, io.Writer, Options) (PieceReceipt, error)
	GetReplica(context.Context, string, Options) (ReplicaCore, error)

	ListFile(context.Context, common.Address, Options) ([]FileReceipt, error)
	ListPiece(context.Context, common.Address, Options) ([]PieceReceipt, error)
	ListReplica(context.Context, common.Address, Options) ([]ReplicaCore, error)

	Request(context.Context, common.Address, string) (common.Address, error)
	Confirm(context.Context, common.Address, ReplicaCore) error
}

type IReplicaStore interface {
	Put(context.Context, PieceCore, []byte) (ReplicaCore, error)
	Get(context.Context, string, io.Writer, Options) (ReplicaCore, error)
	List(context.Context, common.Address, Options) ([]ReplicaCore, error)

	GenProof(context.Context, string, []byte, int) ([]byte, error)
}

type ICommitment interface {
	Add(ICommitment) error
	Serialize() []byte // 48
	Raw() []byte       // 96
}

type IProof interface {
	Type() int
	Add(IProof) error
	Serialize() []byte // 48+48
}

type IChallenge interface {
	Type() int
	Add(ICommitment) error
	Commitment() ICommitment
}

type IPublicKey interface {
	VerifyKey() IVerifyKey

	GenCommitments(int, io.Reader) ([]ICommitment, error)
	GenCommitment(int, []byte, int) (ICommitment, error)
	GenProofs(IChallenge, int, io.Reader) ([]IProof, error)
	GenProof(IChallenge, int, []byte) (IProof, error)

	Serialize() []byte
	Deserialize([]byte) error
}

type IVerifyKey interface {
	VerifyProof(IChallenge, IProof) error

	Serialize() []byte
	Deserialize([]byte) error
}
