package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `grom:"unique"`
	Password string `grom:"size:255" gorm:"not null"`
}
