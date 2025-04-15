package main

import (
	"net/http"

	"github.com/Higor-ViniciusDev/api/configs"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
	"github.com/Higor-ViniciusDev/api/internal/infra/webserver/handles"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	usuarioDB := database.NovoUsuarioDB(db)

	produtctHandler := handles.NovoProdutoHandle(produtoDB)
	usuarioHandler := handles.NovoUsuariohandler(usuarioDB)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/produto/create", produtctHandler.NovoProduto)
	r.Get("/produtos/{id}", produtctHandler.BuscaProduto)
	r.Put("/produto/{id}", produtctHandler.AlteraProduto)
	r.Get("/produtos", produtctHandler.BuscaTodosProdutos)

	r.Post("/usuario/create", usuarioHandler.NovoUsuario)
	http.ListenAndServe(":8080", r)
}
