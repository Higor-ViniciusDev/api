package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPorta       string `mapstructure:"DB_PORT"`
	DBNome        string `mapstructure:"DB_NAME"`
	DBSenha       string `mapstructure:"DB_PASSWORD"`
	DBUsuario     string `mapstructure:"DB_USER"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTTempo      string `mapstructure:"JWT_EXPIRATION"`
	TolkenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	// cfg = &conf{}

	viper.SetConfigName("app.Config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		panic(err)
	}

	//Após lê o arquivo com o unmarshal, ele vai criar o token de autenticação
	// com o algoritmo HS256 e a chave secreta que está no arquivo .env
	cfg.TolkenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
