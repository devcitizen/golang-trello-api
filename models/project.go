package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name   string
	Detail string
	Sprint []Sprint
}
