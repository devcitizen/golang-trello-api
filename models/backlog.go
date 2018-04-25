package models

import (
	"github.com/jinzhu/gorm"
)

type Backlog struct {
	gorm.Model
	Name     string
	Detail   string
	Status   int `gorm:default:0`
	SprintID uint
}
