package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Estoque struct {
	ID                   uuid.UUID
	MaterialCodigo       string
	Material             Material
	QuantidadeDisponivel int16
	EstoqueMinimo        int16
	DataAtualizacao      time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
