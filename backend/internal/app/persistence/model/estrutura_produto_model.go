package model

import "uuid"

type EstruturaProduto struct {
	ID         uuid.UUID
	Produto    Produto
	Componente []Componente
}

type Componente struct {
	ID              uuid.UUID
	Material        Material
	Quantidade      float32
	PerdaPercentual *int8
}
