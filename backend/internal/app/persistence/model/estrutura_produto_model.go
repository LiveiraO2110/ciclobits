package model

import "github.com/google/uuid"

type EstruturaProduto struct {
	ID            uuid.UUID
	ProdutoCodigo string
	Produto       Produto
	Componente    []Componente
}

type Componente struct {
	ID                 uuid.UUID
	EstruturaProdutoID uuid.UUID
	MaterialCodigo     string
	Material           Material
	Quantidade         float32
	PerdaPercentual    *int8
}
