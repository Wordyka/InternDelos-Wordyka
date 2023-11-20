package models

import "gorm.io/gorm"

type Ponds struct {
    gorm.Model
    Name                 string  `gorm:"type:varchar(255); not null" json:"name"`
    FarmID               uint    `gorm:"not null" json:"farmID"`
    Size                 float32 `gorm:"type:float" json:"size"`
    Depth                float32 `gorm:"type:float" json:"depth"`
    TypeOfAquaticSpecies string  `gorm:"type:varchar(100)" json:"typeOfAquaticSpecies"`
    WaterSource          string  `gorm:"type:varchar(100)" json:"waterSource"`
    WaterQuality         string  `gorm:"type:text" json:"waterQuality"`
    Farm                 Farms   `gorm:"foreignkey:FarmID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"farms"`
}