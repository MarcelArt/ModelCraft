package models

import "gorm.io/gorm"

const authorizedDeviceTableName = "authorized_devices"

type AuthorizedDevice struct {
	gorm.Model
	// Insert your fields here
	RefreshToken string `gorm:"not null" json:"-"`
	UserAgent    string `json:"userAgent"`
	Ip           string `json:"ip"`
	UserID       uint   `gorm:"not null" json:"userId"`

	User *User `json:"user"`
}

type AuthorizedDeviceDTO struct {
	DTO
	// Insert your fields here
	RefreshToken string `gorm:"not null" json:"-"`
	UserAgent    string `json:"userAgent"`
	Ip           string `json:"ip"`
	UserID       uint   `gorm:"not null" json:"userId"`
}

type AuthorizedDevicePage struct {
	// Insert your fields here
	UserAgent string `json:"userAgent"`
	Ip        string `json:"ip"`
	UserID    uint   `gorm:"not null" json:"userId"`
}

func (AuthorizedDeviceDTO) TableName() string {
	return authorizedDeviceTableName
}
