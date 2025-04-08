package entity

import (
	"github.com/Higor-ViniciusDev/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    entity.ID `json:"id"`
	Nome  string    `json:"nome"`
	Email string    `json:"email"`
	Senha string    `json:"-"`
}

func NewUser(email, nome, senha string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:    entity.NewID(),
		Nome:  nome,
		Email: email,
		Senha: string(hash),
	}, nil
}

func (u *User) ValidarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))

	return err == nil
}
