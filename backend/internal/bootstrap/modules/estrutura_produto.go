package modules

import "gorm.io/gorm"

type EstruturaProduto struct{}

func NewEstruturaProduto(db *gorm.DB) (EstruturaProduto){
	return EstruturaProduto{}
}