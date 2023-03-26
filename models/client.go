package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Client struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Secret      string `json:"secret"`
	PrivateKey  []byte `json:"-"`
	Name        string `json:"name"`
	RedirectUrl string `json:"redirect_url"`
}

func (c *Client) Save(db *gorm.DB) error {
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&c).Error
	return err
}

func (c *Client) List(db *gorm.DB) (clients []Client, err error) {
	err = db.Model(c).Find(&clients).Error
	return
}

func (c *Client) GetById(userId string, db *gorm.DB) (*Client, error) {
	//addresses := make([]WalletAddress, 0)
	err := db.Model(c).Where("id = ?", userId).Take(c).Error
	return c, err
}
