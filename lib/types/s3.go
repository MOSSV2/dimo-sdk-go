package types

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type BucketInfo struct {
	Name    string
	Deleted bool
	Created time.Time
}

func (bi *BucketInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(bi)
}

func (bi *BucketInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, bi)
}

type ObjectInfo struct {
	Name    string
	ETag    string
	Deleted bool
}

func (oi *ObjectInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(oi)
}

func (oi *ObjectInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, oi)
}

type ListObjectOption struct {
	Prefix, Marker, Delimiter string
	MaxKeys                   int
}

func (oi *ListObjectOption) Serialize() ([]byte, error) {
	return cbor.Marshal(oi)
}

func (oi *ListObjectOption) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, oi)
}

type IObjectStore interface {
	CreateBucket(ctx context.Context, addr common.Address, bucket string) error
	GetBucketInfo(ctx context.Context, addr common.Address, bucket string) (BucketInfo, error)
	DeleteBucket(ctx context.Context, addr common.Address, bucket string) error
	ListBucket(ctx context.Context, addr common.Address) ([]BucketInfo, error)
	PutObject(ctx context.Context, addr common.Address, bucket, object, etag string) error
	GetObjectInfo(ctx context.Context, addr common.Address, bucket, object string) (ObjectInfo, error)
	DeleteObject(ctx context.Context, addr common.Address, bucket, object string) error
	ListObject(ctx context.Context, addr common.Address, bucket string, opt ListObjectOption) (ListObjectResult, error)
}
