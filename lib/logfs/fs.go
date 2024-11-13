package simplefs

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/fxamacker/cbor/v2"
)

var logger = log.Logger("logfs")

var _ types.IFileStore = (*LogFS)(nil)

const (
	MaxSize = 4 * 1024 * 1024
	prefix  = "logfs"
)

type LogMeta struct {
	Index uint32 // which log
	Start uint64
	Size  uint64
	Hash  []byte // 32byte
	Name  []byte
}

func (lm *LogMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(lm)
}

func (lm *LogMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, lm)
}

type LogFS struct {
	sync.RWMutex
	ds       types.IKVStore
	curSize  int64
	curIndex uint32
	curFi    *os.File
	basedir  string
}

func New(ds types.IKVStore, dir string) (*LogFS, error) {
	//log.SetLogLevel("debug")
	logger.Infof("logfs start at: %s", dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	sf := &LogFS{
		basedir: dir,
		ds:      ds,
	}

	dsKey := types.NewKey(types.DsLogFS, prefix)
	val, err := sf.ds.Get(dsKey)
	if err == nil && len(val) == 4 {
		sf.curIndex = binary.BigEndian.Uint32(val)
	}

	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.log", sf.curIndex))
	sf.curFi, err = os.OpenFile(curlog, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	fi, err := sf.curFi.Stat()
	if err != nil {
		return nil, err
	}
	sf.curSize = fi.Size()

	logger.Infof("logfs started at: %s %d %d", dir, sf.curIndex, sf.curSize)
	return sf, nil
}

func (sf *LogFS) forward() error {
	err := sf.curFi.Close()
	if err != nil {
		return err
	}
	sf.curIndex++

	dsKey := types.NewKey(types.DsLogFS, prefix)
	val := make([]byte, 4)
	binary.BigEndian.PutUint32(val, sf.curIndex)
	err = sf.ds.Put(dsKey, val)
	if err != nil {
		return err
	}

	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.log", sf.curIndex))
	fi, err := os.OpenFile(curlog, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	sf.curFi = fi
	sf.curSize = 0
	logger.Infof("logfs forword to: %d", sf.curIndex)
	return nil
}

func (sf *LogFS) Put(key, val []byte) error {
	sum := sha256.Sum256(val)
	dskey := types.NewKey(types.DsLogFS, key)

	sf.Lock()
	defer sf.Unlock()

	has, err := sf.ds.Has(dskey)
	if err == nil && has {
		return fmt.Errorf("already has key: %s", string(key))
	}

	if len(val)+int(sf.curSize) > MaxSize {
		err := sf.forward()
		if err != nil {
			return err
		}
	}

	n, err := sf.curFi.WriteAt(val, sf.curSize)
	if err != nil {
		return err
	}

	logger.Debugf("logfs write at: %d %d %d %d", sf.curIndex, sf.curSize, n, len(val))

	lm := LogMeta{
		Index: sf.curIndex,
		Start: uint64(sf.curSize),
		Size:  uint64(n),
		Hash:  sum[:],
		Name:  key,
	}

	sf.curSize += int64(n)

	lmv, err := lm.Serialize()
	if err != nil {
		return err
	}
	err = sf.ds.Put(dskey, lmv)
	if err != nil {
		return err
	}

	return nil
}

func (sf *LogFS) Get(key []byte, opts ...int) ([]byte, error) {
	dskey := types.NewKey(types.DsLogFS, key)
	val, err := sf.ds.Get(dskey)
	if err != nil {
		return nil, err
	}

	lm := new(LogMeta)
	err = lm.Deserialize(val)
	if err != nil {
		return nil, err
	}

	logger.Infof("logfs read at: %d %d %d", lm.Index, lm.Start, lm.Size)
	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.log", lm.Index))
	fi, err := os.OpenFile(curlog, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	res := make([]byte, lm.Size)
	n, err := fi.ReadAt(res, int64(lm.Start))
	if err != nil {
		return nil, err
	}
	if n != int(lm.Size) {
		return nil, fmt.Errorf("unequal size")
	}

	sum := sha256.Sum256(res)
	if !bytes.Equal(sum[:], lm.Hash) {
		return nil, fmt.Errorf("unequal content")
	}

	return res, nil
}

func (sf *LogFS) Has(key []byte) (bool, error) {
	return false, nil
}

func (sf *LogFS) Delete(key []byte) error {
	return nil
}

func (sf *LogFS) Size() types.DiskStats {
	ds, _ := utils.GetDiskStatus(sf.basedir)
	return ds
}

func (sf *LogFS) Close() error {
	sf.Lock()
	defer sf.Unlock()
	return sf.curFi.Close()
}
