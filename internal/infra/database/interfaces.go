package database

import "github.com/Higor-ViniciusDev/api/internal/entity"

type UsuarioInterface interface {
	NovoUsuario(s *entity.Usuario) error
	ProcuraPorEmail(email string) (*entity.Usuario, error)
}
