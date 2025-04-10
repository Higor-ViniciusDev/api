package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduto(t *testing.T) {
	prod, err := NovoProduto("caixa papelão 12", 64.64)

	assert.Nil(t, err, "Não pode haver error na criacao")
	assert.Empty(t, prod.Nome, "Nome não pode ser vazio")
}
