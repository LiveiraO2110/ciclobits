package modules

import "gorm.io/gorm"

type OrdensProducao struct{}

func NewOrdensProducao(db *gorm.DB) (OrdensProducao){
	return OrdensProducao{}
}