package model

import (
	"time"
	"uuid"
)

type Produto struct {
	Codigo        string `gorm:"primaryKey;size:4;check:chk_produto_codigo,codigo ~ '^P[0-9]{3}$'" json:"codigo"`
	Descricao     string
	UnidadeMedida string
	LoteMinimo    int16
	Ativo         bool
	Roteiro       []RoteiroEtapa
	CratedAt      time.Time
	UpdateAt      time.Time
	DeleteAt      time.Time
}

type RoteiroEtapa struct {
	ID               uuid.UUID
	ProdutoCodigo    string
	Sequencia        int
	Recurso          string
	TempoUnitarioMin float64
	CratedAt         time.Time
	UpdateAt         time.Time
	DeleteAt         time.Time
}


