package models

import "gorm.io/gorm"

type Farms struct {
    gorm.Model
    Name                string  `gorm:"type:varchar(255); not null" json:"name"`
    Location            string  `gorm:"type:varchar(255); not null" json:"location"`
    Owner               string  `gorm:"type:varchar(255)" json:"owner"`
    Size                float32 `gorm:"type:float" json:"size"`
    TypeOfAquaculture   string  `gorm:"type:varchar(100)" json:"typeOfAquaculture"`
    ProductionCapacity  int     `gorm:"type:int" json:"productionCapacity"`
    Ponds               []Ponds `gorm:"foreignKey:FarmID" json:"ponds,omitempty"`
}