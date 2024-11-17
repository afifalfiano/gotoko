package models

import "time"

type Address struct {
	ID         string `gorm:size:36;not null;uniqeIndex;primary_key`
	User       User
	UserId     string `gorm:"size:36;index"`
	Name       string `gorm:"size:36;"`
	IsPrimary  bool
	CityId     string `gorm:"size:255;"`
	ProvinceId string `gorm:"size:255;"`
	Address1   string `gorm:"size:255;"`
	Address2   string `gorm:"size:255;"`
	Phone      string `gorm:"size:100;"`
	Email      string `gorm:"size:100;"`
	PostCode   string `gorm:"size:100;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
