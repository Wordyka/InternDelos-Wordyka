package models

import "gorm.io/gorm"

type Farms struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255); not null" json:"name"`
	Ponds []Ponds `gorm:"foreignKey:FarmID" json:"ponds,omitempty"`
}
	