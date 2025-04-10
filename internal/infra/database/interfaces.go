package database

import "github.com/Higor-ViniciusDev/api/internal/entity"

type UsuarioInterface interface {
	CreateUsuarioDB(s *entity.Usuario) error
	ProcuraPorEmail(email string) (*entity.Usuario, error)
}
