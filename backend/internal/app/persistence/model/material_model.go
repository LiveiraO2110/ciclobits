package model

import (
	"time"
	"uuid"
)

type Material struct {
	Codigo              uuid.UUID
	Descricao           string
	Tipo                string
	UnidadeMedida       string
	PrazoResposicaoDias *int16
	CratedAt            time.Time
	UpdateAt            time.Time
	DeleteAt            time.Time
}