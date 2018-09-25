package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
  gorm.Model
	Name string `gorm:"size:150;unique;not null"valid:"required,length(1|150)"`
}