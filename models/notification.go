package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Notification struct {
  gorm.Model
	Text string `gorm:"size:3000;not null"valid:"required,length(1|3000)"`
	Image string `gorm:"size:150;not null"valid:"required,length(1|150)"`
	CompanyID uint `gorm:"not null"valid:"required"`
}

type NotificationQuery struct {
	Id uint
	CreatedAt time.Time
	Text string
	Image string
	Company string
}

func (dm *DbMethods) FindNotificationById(id uint) *NotificationQuery {
	ntf := &NotificationQuery{}
	dm.DB.Table("notifications n").
		Select(`
			n.id,
			n.created_at,
			n.text,
			n.image,
			c.name company
		`).
		Joins("LEFT JOIN companies c ON n.company_id = c.id").
		Where("n.id = ?", id).
		Order("n.id DESC").
		Limit(1).
		Scan(ntf)
	return ntf
}