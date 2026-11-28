package model

type LinhaProducao struct {
	Codigo        string `gorm:"primaryKey;size:3;check:chk_linha_codigo,codigo GLOB 'L[0-9][0-9]'" json: codigo`
	Nome          string
	CapacidadeDia int16
}
