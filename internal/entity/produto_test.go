package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduto(t *testing.T) {
	prod, err := NovoProduto("caixa papelão 12", 64.64)

	assert.Nil(t, err, "Não pode haver error na criacao")
	assert.NotEmpty(t, prod.Nome, "Nome não pode ser vazio")
	assert.NotEmpty(t, prod.ID, "ID não pode ser vazio")
}
