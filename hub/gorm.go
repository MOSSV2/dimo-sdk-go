package hub

import (
	"os"
	"path/filepath"

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
	db.AutoMigrate(&types.Account{})
	db.AutoMigrate(&types.Needle{})
	db.AutoMigrate(&types.Volume{})
	s.gdb = db
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

func (s *Server) addNeedle(owner, name string, findex uint64, start, length uint64) {
	s.gdb.Create(&types.Needle{
		Owner: owner,
		Name:  name,
		File:  findex,
		Start: start,
		Size:  length,
	})
	logger.Info("create needle: ", owner)
}

func (s *Server) getNeedle(owner, name string) ([]types.Needle, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Name: name, Owner: owner}).Find(&needle)
	if result.Error != nil {
		return needle, result.Error
	}
	return needle, nil
}

func (s *Server) getNeedleDisplay(owner, name string) ([]types.NeedleDisplay, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Name: name, Owner: owner}).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}
	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		vol, err := s.getVolume(needle[i].Owner, needle[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) listNeedle(owner string, offset, limit int) ([]types.Needle, error) {
	var needles []types.Needle
	result := s.gdb.Where(&types.Needle{Owner: owner}).Order("id desc").Limit(limit).Offset(offset).Find(&needles)
	if result.Error != nil {
		return nil, result.Error
	}

	return needles, nil
}

func (s *Server) listNeedleDisplay(owner string, offset, limit int) ([]types.NeedleDisplay, error) {
	logger.Debug("list needle: ", owner, offset, limit)
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Owner: owner}).Order("id desc").Limit(limit).Offset(offset).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}

	res := make([]types.NeedleDisplay, 0, len(needle))
	for i := 0; i < len(needle); i++ {
		nd := types.NeedleDisplay{
			CreatedAt: needle[i].CreatedAt,
			Name:      needle[i].Name,
			Owner:     needle[i].Owner,
			File:      needle[i].File,
			Start:     needle[i].Start,
			Size:      needle[i].Size,
		}
		vol, err := s.getVolume(needle[i].Owner, needle[i].File)
		if err == nil && len(vol) > 0 {
			nd.Piece = vol[0].Piece
			nd.TxHash = vol[0].TxHash
		}
		res = append(res, nd)
	}

	return res, nil
}

func (s *Server) addVolume(owner string, findex uint64, piece, txn string) {
	s.gdb.Create(&types.Volume{
		Owner:  owner,
		File:   findex,
		Piece:  piece,
		TxHash: txn,
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
