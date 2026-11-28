package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	PrioridadeAlta  = "alta"
	PrioridadeMedia = "media"
	PrioridadeBaixa = "baixa"
)

type Pedido struct {
	Numero        string `gorm:"primaryKey"`
	Cliente       string
	ProdutoCodigo string
	Produto       Produto
	Quantidade    int16
	DataPedido    time.Time
	DataEntrega   time.Time
	Prioridade    string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
