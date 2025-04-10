package database

import (
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"gorm.io/gorm"
)

type Usuario struct {
	DB *gorm.DB
}

func (U *Usuario) Criar(s *entity.Usuario) error {
	return U.DB.Create(&s).Error
}

func (U *Usuario) ProcuraPorEmail(email string) (*entity.Usuario, error) {
	var usuario1 entity.Usuario

	if err := U.DB.Where("email = ?", email).First(&usuario1).Error; err != nil {
		return nil, err
	}

	return &usuario1, nil
}
