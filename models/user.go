package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id       int
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     Role   `gorm:"foreignkey:RoleID;PRELOAD:true"`
	RoleID   uint
	Password string `gorm:"type:varchar(100);"`
	Address  string `gorm:"index:addr"` // create index with name `addr` for address
}

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null"`
}

type RoleRaw struct {
	ID   uint `json:"id"`
	Name string
}
