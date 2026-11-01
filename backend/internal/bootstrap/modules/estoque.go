package modules

import "gorm.io/gorm"

type Estoque struct{}

func NewEstoque(db *gorm.DB) (*Estoque){
	return &Estoque{}
}