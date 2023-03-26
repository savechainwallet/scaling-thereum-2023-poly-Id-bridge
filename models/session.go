package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Session struct {
	Id              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId          string `json:"user_id" `
	ClientId        string `json:"client_id"`
	AuthRequest     []byte `json:"auth_request"`
	IsVerified      bool   `json:"is_verified"`
	IsAuthenticated bool   `json:"is_authenticated"`
	Claims          []byte `json:"claims"`
	RequestId       string `json:"request_id"`
	Connected       bool   `json:"connected"`
	Code            string `json:"id_token"`
	RedirectUrl     string `json:"redirect_url"`
}

func (s *Session) Save(db *gorm.DB) error {
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&s).Error
	return err
}

func (s *Session) List(db *gorm.DB) (sessions []Session, err error) {
	err = db.Model(s).Find(&sessions).Error
	return
}

func (s *Session) GetById(id uint, db *gorm.DB) (*Session, error) {
	//addresses := make([]WalletAddress, 0)
	err := db.Model(s).Where("id = ?", id).Take(s).Error
	return s, err
}

func (s *Session) GetByRequestId(requestId string, db *gorm.DB) (*Session, error) {
	//addresses := make([]WalletAddress, 0)
	err := db.Model(s).Where("request_id = ?", requestId).Take(s).Error
	return s, err
}

func (s *Session) GetByCode(code string, db *gorm.DB) (*Session, error) {
	err := db.Model(s).Where("code = ?", code).Take(s).Error
	return s, err

}
