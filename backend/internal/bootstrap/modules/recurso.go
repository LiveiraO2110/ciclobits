package modules

import "gorm.io/gorm"

type Recurso struct{}

func NewRecurso(db *gorm.DB) (*Recurso){
	return &Recurso{}
}