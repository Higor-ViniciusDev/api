package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Higor-ViniciusDev/api/configs"
	"github.com/Higor-ViniciusDev/api/internal/dto"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Produto{}, &entity.Usuario{})

	produtoDB := database.NovoProdutoDB(db)
	produtctHandler := NovoProdutoHandle(produtoDB)

	http.HandleFunc("/produto/create", produtctHandler.NovoProduto)
	http.ListenAndServe(":8080", nil)
}

type ProdutoHandler struct {
	ProdutoDB database.ProdutoInterface
}

func NovoProdutoHandle(db database.ProdutoInterface) *ProdutoHandler {
	return &ProdutoHandler{
		ProdutoDB: db,
	}
}

func (h *ProdutoHandler) NovoProduto(w http.ResponseWriter, r *http.Request) {
	var produtoInput dto.CreateProdutoInput

	err := json.NewDecoder(r.Body).Decode(&produtoInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NovoProduto(produtoInput.Nome, produtoInput.Preco)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProdutoDB.CreateProdutoDB(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Produto Criado com sucesso")
	w.WriteHeader(http.StatusCreated)
}
