package handles

import (
	"encoding/json"
	"net/http"

	"github.com/Higor-ViniciusDev/api/internal/dto"
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/Higor-ViniciusDev/api/internal/infra/database"
)

type UsuarioHandler struct {
	UserDB database.UsuarioInterface
}

func NovoUsuariohandler(db database.UsuarioInterface) *UsuarioHandler {
	return &UsuarioHandler{UserDB: db}
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
