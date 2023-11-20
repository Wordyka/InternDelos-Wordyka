package models

import "gorm.io/gorm"

type Ponds struct {
	gorm.Model
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255); not null" json:"name"`
	FarmID      uint64 `gorm:"not null" json:"-"`
	Farms        Farms   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"farms"`
}
