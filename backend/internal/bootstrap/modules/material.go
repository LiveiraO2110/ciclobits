package modules

import "gorm.io/gorm"

type Material struct{}

func NewMaterial(db *gorm.DB) (Material){
	return Material{}
}