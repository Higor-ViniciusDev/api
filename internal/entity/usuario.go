package entity

import (
	"github.com/Higor-ViniciusDev/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID    entity.ID `json:"id"`
	Nome  string    `json:"nome"`
	Email string    `json:"email"`
	Senha string    `json:"-"`
}

func NovoUsuario(email, nome, senha string) (*Usuario, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &Usuario{
		ID:    entity.NewID(),
		Nome:  nome,
		Email: email,
		Senha: string(hash),
	}

	return user, nil
}

func (u *Usuario) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))

	return err == nil
}
