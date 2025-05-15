package hub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) loadGORM() {
	gpath := filepath.Join(s.rp.Path(), "gorm")

	os.MkdirAll(gpath, os.ModePerm)
	gpath = filepath.Join(gpath, "gorm.db")
	db, err := gorm.Open(sqlite.Open(gpath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto migrate tables
	db.AutoMigrate(&types.Account{})
	db.AutoMigrate(&types.Bucket{})
	db.AutoMigrate(&types.Needle{})
	db.AutoMigrate(&types.Volume{})
	db.AutoMigrate(&types.StatRecord{})
	s.gdb = db

	// iterate all needles to update bucket
	ni := os.Getenv("NEED_INIT")
	if ni != "" {
		var needles []types.Needle
		db.Find(&needles)
		for _, needle := range needles {
			if needle.Bucket != "" {
				continue
			}
			fmt.Println("update bucket: ", needle.Name)
			var w bytes.Buffer
			s.logFSRead(needle.Owner, needle.Name, &w)
			if w.Len() > 0 {
				// decode w to json
				var meta map[string]interface{}
				err := json.Unmarshal(w.Bytes(), &meta)
				if err != nil {
					continue
				}
				bucketName, ok := meta["name"].(string)
				if ok {
					s.addBucket(needle.Owner, bucketName)
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				} else {
					bucketName := needle.Owner
					s.addBucket(needle.Owner, bucketName)
					// update needle
					db.Model(&needle).Update("bucket", bucketName)
				}
			}
		}
	}

}

func (s *Server) addAccount(owner string) {
	var account types.Account
	result := s.gdb.First(&account, "name = ?", owner)
	if result.RowsAffected > 0 {
		logger.Info("already has account: ", owner)
		return
	}

	s.gdb.Create(&types.Account{
		Name: owner,
	})
	logger.Info("create account: ", owner)
}

// TODO: bucket is global unique
func (s *Server) addBucket(owner, bucket string) error {
	var gbucket types.Bucket
	result := s.gdb.First(&gbucket, "name = ? ", bucket)
	if result.RowsAffected > 0 {
		if gbucket.Owner != owner {
			logger.Infof("bucket: %s is owned by %s", bucket, gbucket.Owner)
			return fmt.Errorf("bucket: %s is owned by %s", bucket, gbucket.Owner)
		}
		logger.Info("already has bucket: ", bucket)
		return nil
	}

	s.gdb.Create(&types.Bucket{
		Name:  bucket,
		Owner: owner,
	})
	logger.Info("create bucket: ", bucket)
	return nil
}

func (s *Server) getAccount(owner string) ([]types.Account, error) {
	var accounts []types.Account
	result := s.gdb.Where(&types.Account{Name: owner}).Find(&accounts)
	if result.Error != nil {
		return accounts, result.Error
	}
	return accounts, nil
}

func (s *Server) listAccount(offset, limit int) ([]types.Account, error) {
	var accounts []types.Account
	result := s.gdb.Order("id desc").Limit(limit).Offset(offset).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func (s *Server) getBucket(owner, bucket string) ([]types.Bucket, error) {
	var buckets []types.Bucket
	result := s.gdb.Where(&types.Bucket{Owner: owner, Name: bucket}).Find(&buckets)
	if result.Error != nil {
		return nil, result.Error
	}
	return buckets, nil
}

func (s *Server) listBucket(owner string, offset, limit int) ([]types.Bucket, error) {
	var buckets []types.Bucket
	result := s.gdb.Where(&types.Bucket{Owner: owner}).Order("id desc").Limit(limit).Offset(offset).Find(&buckets)
	if result.Error != nil {
		return nil, result.Error
	}

	return buckets, nil
}

func (s *Server) addNeedle(owner, bucket, name string, findex uint64, start, length uint64) {
	s.gdb.Create(&types.Needle{
		Owner:  owner,
		Bucket: bucket,
		Name:   name,
		File:   findex,
		Start:  start,
		Size:   length,
	})
	logger.Info("create needle: ", owner)
}

func (s *Server) getNeedleByName(name string) ([]types.Needle, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Name: name}).Find(&needle)
	if result.Error != nil {
		return needle, result.Error
	}
	return needle, nil
}

func (s *Server) getNeedleDisplay(owner, bucket, name string) ([]types.NeedleDisplay, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Name: name, Owner: owner, Bucket: bucket}).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}
	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			Bucket:    needle[i].Bucket,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		vol, err := s.getVolume(needle[i].Owner, needle[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
			nd.ChainType = vol[0].ChainType
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) listNeedle(owner, bucket string, offset, limit int) ([]types.Needle, error) {
	var needles []types.Needle
	result := s.gdb.Where(&types.Needle{Owner: owner, Bucket: bucket}).Order("id desc").Limit(limit).Offset(offset).Find(&needles)
	if result.Error != nil {
		return nil, result.Error
	}

	return needles, nil
}

func (s *Server) listNeedleDisplay(owner, bucket string, offset, limit int) ([]types.NeedleDisplay, error) {
	logger.Debug("list needle: ", owner, bucket, offset, limit)
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Owner: owner, Bucket: bucket}).Order("id desc").Limit(limit).Offset(offset).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}

	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			Bucket:    needle[i].Bucket,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		vol, err := s.getVolume(needle[i].Owner, needle[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
			nd.ChainType = vol[0].ChainType
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) addVolume(owner string, findex uint64, piece, txn string) {
	s.gdb.Create(&types.Volume{
		ChainType: s.rp.Repo().Config().Chain.Type,
		Owner:     owner,
		File:      findex,
		Piece:     piece,
		TxHash:    txn,
	})
	logger.Info("create volume: ", piece)
}

func (s *Server) getVolume(owner string, fid uint64) ([]types.Volume, error) {
	var vol []types.Volume
	result := s.gdb.Where(&types.Volume{Owner: owner, File: fid}).Find(&vol)
	if result.Error != nil {
		return vol, result.Error
	}
	return vol, nil
}

func (s *Server) listVolume(owner string, offset, limit int) ([]types.Volume, error) {
	var vols []types.Volume
	result := s.gdb.Where(&types.Volume{Owner: owner}).Order("id desc").Limit(limit).Offset(offset).Find(&vols)
	if result.Error != nil {
		return nil, result.Error
	}

	return vols, nil
}

func (s *Server) listConversation(ctx context.Context, addr, bucket string) ([]string, error) {
	var needles []types.Needle
	// create time is time.Time >= 2025-03-07,
	// name end with "_0"
	// addr may be empty
	// bucket may be empty
	var result *gorm.DB
	if bucket == "" {
		result = s.gdb.Model(&types.Needle{}).Where("owner = ? and created_at >= ? and name like ?", addr, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC), "%_0").Find(&needles)
	} else {
		result = s.gdb.Model(&types.Needle{}).Where("owner = ? and bucket = ? and created_at >= ? and name like ?", addr, bucket, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC), "%_0").Find(&needles)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	// name is conversation_index, so we need to split it to conversation_id
	// reomve duplicate
	conversations := make([]string, 0, len(needles))
	cmap := make(map[string]struct{})
	for _, needle := range needles {
		if len(needle.Name) <= 2 {
			continue
		}
		parts := strings.Split(needle.Name, "_")
		if len(parts) != 2 {
			continue
		}
		conversationID := parts[0]
		if _, ok := cmap[conversationID]; !ok {
			conversations = append(conversations, conversationID)
			cmap[conversationID] = struct{}{}
		}
	}
	return conversations, nil
}

func (s *Server) getConversation(ctx context.Context, conversation, addr, bucket string) ([]string, error) {
	var needles []types.Needle
	// name contains conversation + "_"
	// create time >= 2025-03-07
	// order by id asc
	var result *gorm.DB
	if bucket == "" {
		result = s.gdb.Model(&types.Needle{}).Where("name like ? and owner = ? and created_at >= ?", conversation+"_%", addr, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)).Order("id asc").Find(&needles)
	} else {
		result = s.gdb.Model(&types.Needle{}).Where("name like ? and owner = ? and bucket = ? and created_at >= ?", conversation+"_%", addr, bucket, time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)).Order("id asc").Find(&needles)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	// name is conversation_index, so we need to split it to conversation_id
	// reomve duplicate
	conversations := make([]string, 0, len(needles))
	for _, needle := range needles {
		var w bytes.Buffer
		_, err := s.download(ctx, needle.Name, addr, &w)
		if err != nil {
			continue
		}
		conversations = append(conversations, w.String())
	}
	return conversations, nil
}
