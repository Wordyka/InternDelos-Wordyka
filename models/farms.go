package models

import "gorm.io/gorm"

type Farms struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
}
