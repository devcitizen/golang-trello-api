package models

import (
	"github.com/jinzhu/gorm"
)

type Sprint struct {
	gorm.Model
	Name      string
	ProjectID uint
	Backlog   []Backlog
}
