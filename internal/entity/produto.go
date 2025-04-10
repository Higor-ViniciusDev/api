package entity

import (
	"errors"
	"time"

	"github.com/Higor-ViniciusDev/api/pkg/entity"
)

var (
	ErrorIdRequerido    = errors.New("id é requerido")
	ErrorIdInvalido     = errors.New("id deve ser um ID válido")
	ErrorNomeRequerido  = errors.New("nome é requerido")
	ErrorPrecoRequerido = errors.New("preco é requerido")
	ErrorPrecoInvalido  = errors.New("preco deve ser maior que 0")
)

type Produto struct {
	ID        entity.ID `json:"id"`
	Nome      string    `json:"nome"`
	Preco     float64   `json:"preco"`
	CreatedAt time.Time `json:"created_at"`
}

func NovoProduto(nome string, preco float64) (*Produto, error) {
	prod := &Produto{
		ID:        entity.NewID(),
		Nome:      nome,
		Preco:     preco,
		CreatedAt: time.Now(),
	}

	if prod.Validar() != nil {
		return nil, prod.Validar()
	}

	return prod, nil
}

func (p *Produto) Validar() error {
	if p.ID.String() == "" {
		return ErrorIdRequerido
	}

	if _, err := entity.PaserID(p.ID.String()); err != nil {
		return ErrorIdInvalido
	}

	if p.Nome == "" {
		return ErrorNomeRequerido
	}
	if p.Preco <= 0 {
		return ErrorPrecoInvalido
	}

	return nil
}
