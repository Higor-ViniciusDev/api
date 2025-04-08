package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNovoUsuario(t *testing.T) {
	user, err := NewUser("teste@gmail.com", "Higor", "123456")

	assert.Nil(t, err, "Não pode houver erro na criação do usuario")
	assert.NotEmpty(t, user.ID, "ID não pode ser vazio")
	assert.NotEmpty(t, user.Nome, "Nome não pode ser vazio")
	assert.NotEmpty(t, user.Email, "Email não pode ser vazio")
	assert.NotEmpty(t, user.Senha, "Senha não pode ser vazia")
}

func TestValidarSenha(t *testing.T) {
	user, _ := NewUser("teste@gmail.com", "Higor", "123456")

	assert.True(t, user.ValidarSenha("123456"), "Senha deve ser válida")
	assert.False(t, user.ValidarSenha("1234567"), "Senha deve ser inválida")
}
