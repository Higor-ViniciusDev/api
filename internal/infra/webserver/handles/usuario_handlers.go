package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Higor-ViniciusDev/api/internal/dto"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UsuarioHandler struct {
	UserDB   database.UsuarioInterface
	Jwt      *jwtauth.JWTAuth
	JwtTempo int
}

func NovoUsuariohandler(db database.UsuarioInterface, jwt *jwtauth.JWTAuth, tempoExpiracao int) *UsuarioHandler {
	return &UsuarioHandler{
		UserDB:   db,
		Jwt:      jwt,
		JwtTempo: tempoExpiracao,
	}
}

func (h *UsuarioHandler) NovoUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioInput dto.CreateUsuarioInpunt

	err := json.NewDecoder(r.Body).Decode(&usuarioInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NovoUsuario(usuarioInput.Email, usuarioInput.Nome, usuarioInput.Senha)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.UserDB.CreateUsuarioDB(u)
	w.WriteHeader(http.StatusOK)
}

func (h *UsuarioHandler) PegaJWT(w http.ResponseWriter, r *http.Request) {
	var JWTDto dto.GetJWT

	err := json.NewDecoder(r.Body).Decode(&JWTDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.ProcuraPorEmail(JWTDto.Email)
	fmt.Println(u)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidarSenha(JWTDto.Senha) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tolkenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtTempo)).Unix(),
	})
	tolken := struct {
		AccessTolken string `json:"access_tolken"`
	}{
		AccessTolken: tolkenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tolken)
}
