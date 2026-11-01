package modules

import "gorm.io/gorm"

type Produto struct{}

func NewProduto(db *gorm.DB) (Produto){
	return Produto{}
}