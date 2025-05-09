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
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

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
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

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

func TestProcuraPorID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Produto{})
	produto, err := entity.NovoProduto("Produto", rand.Float64()*100)
	assert.Nil(t, err)
	newProductDb := NovoProdutoDB(db)

	err = newProductDb.CreateProdutoDB(produto)
	assert.Nil(t, err)

	novoProduto, err := newProductDb.ProcuraPorID(produto.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, produto.ID.String(), novoProduto.ID.String())
}

func TestAlteraProduto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}
	//Cria as tabela
	db.AutoMigrate(&entity.Produto{})
	produto, err := entity.NovoProduto("Produto", rand.Float64()*100)
	assert.Nil(t, err)

	//Cria conexão
	newProductDb := NovoProdutoDB(db)
	db.Create(&produto)

	//Alterar produto
	produto.Nome = "NOVO PRODUTO"
	produto.Preco = 100.0

	//chama a função de alteração
	err = newProductDb.AlteraProduto(produto)
	assert.Nil(t, err, "Não pode haver erro na alteração")

	var produtoAlterado entity.Produto
	err = db.First(&produtoAlterado, "id = ?", produto.ID).Error
	//Validando de tudo deu certo
	assert.Nil(t, err)
	assert.Equal(t, "NOVO PRODUTO", produtoAlterado.Nome)
	assert.Equal(t, 100.0, produtoAlterado.Preco)

	//Validar erro se passar uma entidade não criada
	novoProd, _ := entity.NovoProduto("Produto", rand.Float64()*100)
	err = newProductDb.AlteraProduto(novoProd)
	assert.NotNil(t, err)
}

func TestDeletaProduto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}
	//Cria as tabela
	db.AutoMigrate(&entity.Produto{})
	produto, err := entity.NovoProduto("Produto", rand.Float64()*100)
	assert.Nil(t, err)

	//Cria conexão
	newProductDb := NovoProdutoDB(db)
	db.Create(&produto)

	err = newProductDb.Apagar(produto.ID.String())
	assert.Nil(t, err)

	var produtoDeletado entity.Produto
	err = db.First(&produtoDeletado, "id = ?", produto.ID).Error
	assert.NotNil(t, err)
	assert.Empty(t, produtoDeletado)
}
