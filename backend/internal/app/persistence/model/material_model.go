package model

import (
	"time"

	"gorm.io/gorm"
)

type Material struct {
	Codigo              string `gorm:"primaryKey;size:4;check:chk_material_codigo,codigo GLOB 'M[0-9][0-9][0-9]'" json:codigo`
	Descricao           string
	Tipo                string
	UnidadeMedida       string
	PrazoResposicaoDias *int16
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
