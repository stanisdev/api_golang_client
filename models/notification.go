package models

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
  gorm.Model
	Text string `gorm:"size:3000;not null"valid:"required,length(1|3000)"`
	Image string `gorm:"size:150;unique;not null"valid:"required,length(1|150)"`
	Company string `gorm:"size:150;unique;not null"valid:"required,length(1|150)"`
}