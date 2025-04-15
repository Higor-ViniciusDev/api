package dto

type CreateProdutoInput struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

type CreateUsuarioInpunt struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type GetJWT struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
