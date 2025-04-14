package dto

type CreateProdutoInput struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
