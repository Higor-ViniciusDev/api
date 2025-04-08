package main

import "github.com/Higor-ViniciusDev/api/configs"

func main() {
	cfg, _ := configs.LoadConfig("./")

	println("Servidor rodando na porta: ", cfg.WebServerPort)
}
