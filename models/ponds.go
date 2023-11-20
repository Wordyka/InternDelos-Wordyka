package models

import "gorm.io/gorm"

type Ponds struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255); not null" json:"name"`
	FarmID uint   `gorm:"not null" json:"-"`
	Farm   Farms  `gorm:"foreignkey:FarmID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"farms"`		
}
