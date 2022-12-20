package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `grom:"type:varchar(255);not null"`
	Name         string `gorm:"type:varchar(255);not null"`
	LastName     string `gorm:"type:varchar(255);not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	Age          uint8
	Birthday     string `gorm:"type:varchar(255);not null"`
	RefreshToken string `gorm:"type:varchar(255);not null"`
	Image        string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:varchar(300)"`
	Online       bool   `gorm:"type:bool;default:false"`
	LastSeen     string `gorm:"type:varchar(255)"`
}
