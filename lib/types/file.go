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
	N, K uint8
}

type FileCore struct {
	Policy   Policy
	Name     string
	Hash     string
	Size     int64
	Owner    common.Address
	Append   bool
	Creation time.Time
}

func (frc *FileCore) Serialize() ([]byte, error) {
	return cbor.Marshal(frc)
}

func (frc *FileCore) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, frc)
}

type FileReceipt struct {
	FileCore
	Pieces []string
}

func (fr *FileReceipt) Serialize() ([]byte, error) {
	return cbor.Marshal(fr)
}

func (fr *FileReceipt) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, fr)
}

type PieceCore struct {
	Policy Policy
	Name   string
	Serial uint64
	Size   int64 // raw size
	Start  uint64
	Expire uint64
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
	Serial   uint64
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
	AddFile(context.Context, FileReceipt) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
	GetPiece(context.Context, string, Options) (PieceReceipt, error)
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaCore, error)

	ListFile(context.Context, common.Address, Options) ([]FileReceipt, error)
	ListPiece(context.Context, common.Address, Options) ([]PieceReceipt, error)
	ListReplica(context.Context, common.Address, Options) ([]ReplicaCore, error)
}

type IReplicaStore interface {
	Put(context.Context, PieceCore, []byte) (ReplicaCore, error)
	Get(context.Context, string, io.Writer, Options) (ReplicaCore, error)
	List(context.Context, common.Address, Options) ([]ReplicaCore, error)

	GenProof(context.Context, string, []byte, int) ([]byte, error)
}

type IPieceStore interface {
	PutPiece(context.Context, PieceCore, bool) error
	GetPiece(context.Context, string, io.Writer, Options) (PieceReceipt, error)

	PutReplica(context.Context, ReplicaCore, []byte, bool) error
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaCore, error)

	PutFile(context.Context, FileCore, []string, bool) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
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
