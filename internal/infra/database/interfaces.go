package database

import "github.com/Higor-ViniciusDev/api/internal/entity"

type UsuarioInterface interface {
	CreateUsuarioDB(s *entity.Usuario) error
	ProcuraPorEmail(email string) (*entity.Usuario, error)
}

type ProdutoInterface interface {
	CreateProdutoDB(p *entity.Produto) error
	AlteraProduto(p *entity.Produto) error
	ProcuraPorID(id string) (*entity.Produto, error)
	ProcuraTodos() ([]entity.Produto, error)
	Apagar(id string) error
}
