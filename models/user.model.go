package models

import (
	"time"

	"github.com/MarcelArt/ModelCraft/enums"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

const userTableName = "users"

type User struct {
	gorm.Model
	// Insert your fields here
	Username   string    `json:"username" gorm:"unique;not null"`
	Email      string    `json:"email" gorm:"unique;not null" validate:"email"`
	Password   string    `json:"-" gorm:"not null" validate:"min=8"`
	Salt       string    `json:"-" gorm:"not null"`
	VerifiedAt time.Time `json:"verifiedAt"`
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

func (m UserDTO) AccessClaims(exp int64) jwt.MapClaims {
	return jwt.MapClaims{
		"username": m.Username,
		"userId":   m.ID,
		"exp":      exp,
	}
}

func (m UserDTO) RefreshClaims(isRemember bool) jwt.MapClaims {
	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}
	return jwt.MapClaims{
		"userId":     m.ID,
		"isRemember": isRemember,
		"exp":        expireAt.Unix(),
	}
}
