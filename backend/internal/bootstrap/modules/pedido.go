package modules

import "gorm.io/gorm"

type Pedido struct{}

func NewPedido(db *gorm.DB) (Pedido){
	return Pedido{}
}