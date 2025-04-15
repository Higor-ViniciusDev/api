package main

import (
	"log"
	"net/http"

	"github.com/Higor-ViniciusDev/api/configs"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
	"github.com/Higor-ViniciusDev/api/internal/infra/webserver/handles"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig("./")
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
	usuarioHandler := handles.NovoUsuariohandler(usuarioDB, configs.TolkenAuth, configs.JWTTempo)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/produtos", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TolkenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/create", produtctHandler.NovoProduto)
		r.Get("/{id}", produtctHandler.BuscaProduto)
		r.Put("/{id}", produtctHandler.AlteraProduto)
		r.Get("/", produtctHandler.BuscaTodosProdutos)
	})

	r.Post("/usuario/create", usuarioHandler.NovoUsuario)
	r.Post("/usuario/generateTolken", usuarioHandler.PegaJWT)
	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Passei aqui sei la, processei algo, %v", r.URL)
		next.ServeHTTP(w, r)
	})
}
