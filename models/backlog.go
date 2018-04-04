package models

import (
	"github.com/jinzhu/gorm"
)

type Backlog struct {
	gorm.Model
	Name     string
	Detail   string
	Sprint   Sprint `gorm:"foreignkey:SprintID;PRELOAD:true";nullable`
	SprintID uint
}
