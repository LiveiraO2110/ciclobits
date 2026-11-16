package model

import "time"

type OrdemProducao struct {
	Numero              string
	Produto             Produto
	Pedido              *Pedido
	Prioridade          string
	SequenciaProgamacao int16
	QuantidadePlanejada int16
	QuantidadeProduzida *int16
	DataEmissao         time.Time
	DataInicioPrevista  time.Time
	DataFimPrevista     time.Time
	DataConclusao       *time.Time
	Status              string
	Observacao          *string
	CratedAt            time.Time
	UpdateAt            time.Time
	DeleteAt            time.Time
}

const (
	OrdemProducaoPlanejada = "planejada" 
	OrdemProducaoLiberada = "liberada"
	OrdemProducaoEmProducao = "em_producao"
	OrdemProducaoCancelada = "cancelada"
	OrdemProducaoConcluida = "concluida"
)
