package model

import (
	"time"

	"gorm.io/gorm"
)

type Recurso struct {
	Codigo             string `gorm:"primaryKey"`
	Descricao          string
	Tipo               string
	CapacidadeHorasDia int16 `gorm:"check:chk_recurso_capacidade,capacidade_horas_dia > 0 AND capacidade_horas_dia <= 24"`
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
