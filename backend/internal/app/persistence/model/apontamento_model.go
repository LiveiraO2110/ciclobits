package model

import (
	"time"

	"github.com/google/uuid"
)

type Apontamento struct {
	ID                  uuid.UUID
	OrdemProducaoNumero string
	OrdemProducao       OrdemProducao
	Quantidade          int16
	DataApontamento     time.Time
	Observacao          *string
}
