package model

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func novoDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&Produto{}, &RoteiroEtapa{}, &Recurso{}); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestProdutoCodigoInvalido(t *testing.T) {
	db := novoDB(t)
	err := db.Create(&Produto{Codigo: "X12"}).Error

	if err == nil {
		t.Fatal("esperava erro do CHECK para código fora do padrão P000")
	}
}
