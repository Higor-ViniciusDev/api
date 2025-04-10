package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNovoProdutoBanco(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Produto{})
	prod, _ := entity.NovoProduto("caixa papelão 12", 64.64)
	prodDB := NovoProdutoDB(db)

	err = prodDB.CreateProdutoDB(prod)
	assert.Nil(t, err, "Não pode haver erro na criação do produto no banco")

	var produtoProcurado entity.Produto

	err = db.First(&produtoProcurado, "id = ?", prod.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, prod.Nome, produtoProcurado.Nome)
	assert.Equal(t, prod.Preco, produtoProcurado.Preco)
}

func TestProcuraTodos(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Produto{})
	newProductDb := NovoProdutoDB(db)
	for i := 1; i < 24; i++ {
		produto, err := entity.NovoProduto(fmt.Sprintf("Produto %v", i), rand.Float64()*100)
		assert.Nil(t, err)

		db.Create(produto)
	}

	produtos, err := newProductDb.ProcuraTodos(10, 1, "ASC")
	assert.Nil(t, err)
	assert.Equal(t, "Produto 10", produtos[9].Nome)

	produtos, err = newProductDb.ProcuraTodos(10, 2, "ASC")
	assert.Nil(t, err)
	assert.Equal(t, "Produto 20", produtos[9].Nome)
}
