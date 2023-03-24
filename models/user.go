package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	Id         string `gorm:"primaryKey" json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsAccepted bool   `json:"is_accepted"`
	Country    string `json:"country"`
	Email      string `json:"email"`
}

func (u *User) Save(db *gorm.DB) error {
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&u).Error
	return err
}

func (u *User) List(db *gorm.DB) (users []User, err error) {
	err = db.Model(u).Find(&users).Error
	return
}

func (u *User) GetById(userId string, db *gorm.DB) (*User, error) {
	//addresses := make([]WalletAddress, 0)
	err := db.Model(u).Where("id = ?", userId).Take(u).Error
	return u, err
}
