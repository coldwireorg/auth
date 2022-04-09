package models

import (
	"auth/database"
)

type User struct {
	Name       string `gorm:"not null;unique;primaryKey"`
	Password   string `gorm:"not null"`
	Group      string `gorm:"not null"`
	PrivateKey string `gorm:"not null"`
	PublicKey  string `gorm:"not null"`
}

func (u User) Create() error {
	return database.DB.Create(&u).Error
}

func (u User) Delete() error {
	return database.DB.Delete(&u).Error
}

func (u User) Find() (User, error) {
	var usr User
	err := database.DB.Find(&usr, u).Error
	if err != nil {
		return User{}, err
	}

	return usr, nil
}

func (u User) Exist() bool {
	var usr User
	err := database.DB.Find(&usr, u).Error
	if err != nil {
		return false
	}

	if usr.Name != "" {
		return true
	}

	return false
}

func (u User) IsFirstOne() bool {
	var usr User
	err := database.DB.Limit(1).Find(&usr).Error
	if err != nil {
		return false
	}

	if usr.Name == "" {
		return true
	}

	return false
}

func (u User) Pubkey() (string, error) {
	var pubkey string

	err := database.DB.Select("PublicKey").Find(&u).Scan(&pubkey).Error
	if err != nil {
		return "", err
	}

	return pubkey, nil
}

func (u User) SetPassword(new string) error {
	err := database.DB.Model(User{}).Where("name = ?", u.Name).Update("password", new).Error
	if err != nil {
		return err
	}

	return nil
}
