package model

import (
	"time"

	"gorm.io/gorm"
)

type OrdemProducao struct {
	Numero               string `gorm:"primaryKey"`
	LinhaCodigo          string
	LinhaProducao        LinhaProducao `gorm:"foreignKey:LinhaCodigo"`
	ProdutoCodigo        string
	Produto              Produto
	PedidoNumero         *string
	Pedido               *Pedido
	Prioridade           string
	SequenciaProgramacao int16
	QuantidadePlanejada  int16
	QuantidadeProduzida  *int16
	DataEmissao          time.Time
	DataInicioPrevista   time.Time
	DataFimPrevista      time.Time
	DataConclusao        *time.Time
	Status               string
	Observacao           *string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

const (
	OrdemProducaoPlanejada  = "planejada"
	OrdemProducaoLiberada   = "liberada"
	OrdemProducaoEmProducao = "em_producao"
	OrdemProducaoCancelada  = "cancelada"
	OrdemProducaoConcluida  = "concluida"
)
