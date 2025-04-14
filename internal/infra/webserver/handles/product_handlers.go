package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Higor-ViniciusDev/api/internal/dto"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
	pkgEntity "github.com/Higor-ViniciusDev/api/pkg/entity"
	"github.com/go-chi/chi/v5"
)

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

func (h *ProdutoHandler) BuscaProduto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	produto, err := h.ProdutoDB.ProcuraPorID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produto)
}

func (h *ProdutoHandler) AlteraProduto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var produto entity.Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	produto.ID, err = pkgEntity.PaserID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProdutoDB.AlteraProduto(&produto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println("Produto alterado com sucesso")
	w.WriteHeader(http.StatusOK)
}
