package models

import "gorm.io/gorm"

const userTableName = "users"

type User struct {
	gorm.Model
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null" validate:"email"`
	Password string `json:"-" gorm:"not null" validate:"min=8"`
	Salt     string `json:"-" gorm:"not null"`
}

type UserDTO struct {
	DTO
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null" validate:"email"`
	Password string `json:"password" gorm:"not null" validate:"min=8"`
	Salt     string `json:"-" gorm:"not null"`
}

type UserPage struct {
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null" validate:"email"`
}

type LoginInput struct {
	Username   string `json:"username" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null" validate:"min=8"`
	IsRemember bool   `json:"isRemember"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshInput struct {
	RefreshToken string `json:"refreshToken" validate:"jwt"`
}

func (UserDTO) TableName() string {
	return userTableName
}
