package models

import (
	"github.com/jinzhu/gorm"
)

type Sprint struct {
	gorm.Model
	Name      string `json:"name" validate:"nonzero" gorm:"not null"`
	ProjectID uint   `json:"project_id" validate:"nonzero" gorm:"not null"`
	Backlog   []Backlog
}
