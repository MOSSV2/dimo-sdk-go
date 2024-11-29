package hub

import "github.com/MOSSV2/dimo-sdk-go/lib/types"

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

func (s *Server) listAccountInGORM() ([]types.Account, error) {
	var accounts []types.Account
	result := s.gdb.Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}

func (s *Server) addNeedle(owner, name string, findex uint32, start, length uint64) {
	s.gdb.Create(&types.Needle{
		Owner: owner,
		Name:  name,
		File:  findex,
		Start: start,
		Size:  length,
	})
	logger.Info("create needle: ", owner)
}

func (s *Server) getNeedle(owner, name string) {
	var needle types.Needle
	result := s.gdb.Where(&types.Needle{Name: name, Owner: owner}).Last(&needle)
	if result.Error != nil {
		return
	}
	logger.Info("create needle: ", owner)
}

func (s *Server) listNeedle(owner string, offset, limit int) ([]types.Needle, error) {
	var needle []types.Needle
	result := s.gdb.Where(&types.Needle{Owner: owner}).Order("id desc").Limit(limit).Offset(offset).Find(&needle)
	if result.Error != nil {
		return nil, result.Error
	}

	return needle, nil
}
