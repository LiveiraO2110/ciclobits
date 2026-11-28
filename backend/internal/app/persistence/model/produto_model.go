package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Produto struct {
	Codigo     string `gorm:"primaryKey;size:4;check:chk_produto_codigo,codigo GLOB 'P[0-9][0-9][0-9]'" json:"codigo"`
	Descricao  string
	LoteMinimo int16
	Ativo      bool
	Roteiro    []RoteiroEtapa
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type RoteiroEtapa struct {
	ID               uuid.UUID
	ProdutoCodigo    string
	Produto          Produto
	Sequencia        int
	RecursoCodigo    string
	Recurso          Recurso
	TempoUnitarioMin float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
