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

type Error struct {
	Mensagem string `json:"menssagem"`
}

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

// Novo usuario godoc
// @Sumary NovoUsuario usuario
// @Description Cria novo usuario
// @Tags usuario
// @Accept json
// @Produce json
// @Param request body dto.CreateUsuarioInpunt true "usuario request"
// @Success 201
// @Failure 500  {object} Error
// @Router /usuario/create [post]
func (h *UsuarioHandler) NovoUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioInput dto.CreateUsuarioInpunt

	err := json.NewDecoder(r.Body).Decode(&usuarioInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		erros := Error{Mensagem: err.Error()}
		json.NewEncoder(w).Encode(erros)
		return
	}

	u, err := entity.NovoUsuario(usuarioInput.Email, usuarioInput.Nome, usuarioInput.Senha)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		erros := Error{Mensagem: err.Error()}
		json.NewEncoder(w).Encode(erros)
		return
	}

	h.UserDB.CreateUsuarioDB(u)
	w.WriteHeader(http.StatusOK)
}

// Gerar Novo JWT godoc
// @Sumary Gerar Novo Tolken de acesso
// @Description Gerar Novo Tolken de acesso
// @Tags usuario
// @Accept json
// @Produce json
// @Param request body dto.GetJWT true "usuario credentials"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 404  {object} Error
// @Failure 500  {object} Error
// @Router /usuario/generateTolken [post]
func (h *UsuarioHandler) PegaJWT(w http.ResponseWriter, r *http.Request) {
	var JWTDto dto.GetJWT

	err := json.NewDecoder(r.Body).Decode(&JWTDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		erros := Error{Mensagem: err.Error()}
		json.NewEncoder(w).Encode(erros)
		return
	}

	u, err := h.UserDB.ProcuraPorEmail(JWTDto.Email)
	fmt.Println(u)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		erros := Error{Mensagem: err.Error()}
		json.NewEncoder(w).Encode(erros)
		return
	}

	if !u.ValidarSenha(JWTDto.Senha) {
		w.WriteHeader(http.StatusUnauthorized)
		erros := Error{Mensagem: "Senha incorreta"}
		json.NewEncoder(w).Encode(erros)
		return
	}

	_, tolkenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtTempo)).Unix(),
	})

	tolken := dto.GetJWTOutput{AccessTolken: tolkenString}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tolken)
}
