package model

import "time"

type Recurso struct {
	Codigo             string
	Descricao          string
	Tipo               string
	CapacidadeHorasDia int16
	Status             string
	CratedAt           time.Time
	UpdateAt           time.Time
	DeleteAt           time.Time
}
