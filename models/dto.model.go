package models

import "time"

type IDTO interface {
	GetID() uint
}

type DTO struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (d DTO) GetID() uint {
	return d.ID
}
