package model

import (
	"time"
	"uuid"
)

type Estoque struct {
	ID                   uuid.UUID
	Material             Material
	QuantidadeDisponivel int16
	EstoqueMinimo        int16
	DataAtualizacao      time.Time
	CratedAt             time.Time
	UpdateAt             time.Time
	DeleteAt             time.Time
}
