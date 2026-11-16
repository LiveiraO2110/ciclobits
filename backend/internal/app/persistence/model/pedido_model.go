package model

import "time"

const (
	PrioridadeAlta  = "alta"
	PrioridadeMedia = "media"
	PrioridadeBaixa = "baixa"
)

type Pedido struct {
	Numero      string
	Cliente     string
	Produto     Produto
	Quantidade  int16
	DataPedido  time.Time
	DataEntrega time.Time
	Prioridade  string
	Status      string
	CratedAt    time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
}