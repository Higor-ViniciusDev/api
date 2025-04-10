package database

import (
	"testing"

	"github.com/Higor-ViniciusDev/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNovoUsuarioBanco(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Usuario{})
	user, _ := entity.NovoUsuario("higor@gmail.com", "higor", "123456")
	userDb := NovoUsuarioDB(db)

	err = userDb.CreateUsuarioDB(user)
	assert.Nil(t, err, "Não pode haver erro na criação do usuario no banco")

	var usuarioProcurado entity.Usuario

	err = db.First(&usuarioProcurado, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.Email, usuarioProcurado.Email)
	assert.Equal(t, user.Nome, usuarioProcurado.Nome)
	assert.NotEmpty(t, usuarioProcurado.Senha)
}

func TestProcuraPorEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Usuario{})
	user, _ := entity.NovoUsuario("higor@gmail.com", "higor", "123456")
	userDb := NovoUsuarioDB(db)

	db.AutoMigrate(&entity.Usuario{})

	userDb.CreateUsuarioDB(user)
	usuarioAchado, err := userDb.ProcuraPorEmail("higor@gmail.com")
	assert.Nil(t, err, "Usuario tem que ser achado")
	assert.Equal(t, user.Email, usuarioAchado.Email)
	assert.Equal(t, user.Nome, usuarioAchado.Nome)
	assert.NotEmpty(t, usuarioAchado.Senha)

	usuarioAchado, err = userDb.ProcuraPorEmail("higor123@gmail.com")
	assert.Nil(t, usuarioAchado)
	assert.NotNil(t, err)
}
