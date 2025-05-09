package database

import (
	"github.com/Higor-ViniciusDev/api/internal/entity"
	"gorm.io/gorm"
)

type Produto struct {
	DB *gorm.DB
}

func NovoProdutoDB(db *gorm.DB) *Produto {
	return &Produto{DB: db}
}

// Find All com paginação
func (pr *Produto) ProcuraTodos(limite, pagina int, ordem string) ([]entity.Produto, error) {
	var products []entity.Produto
	var err error

	if ordem != "" && ordem != "asc" && ordem != "desc" {
		ordem = "asc"
	}

	if limite != 0 && pagina != 0 {
		err = pr.DB.Limit(limite).Offset((pagina - 1) * limite).Order("created_at " + ordem).Find(&products).Error
	} else {
		err = pr.DB.Order("created_at " + ordem).Find(&products).Error
	}

	return products, err
}

func (pr *Produto) CreateProdutoDB(p *entity.Produto) error {
	return pr.DB.Create(&p).Error
}

func (pr *Produto) AlteraProduto(p *entity.Produto) error {
	_, err := pr.ProcuraPorID(p.ID.String())

	if err != nil {
		return err
	}

	return pr.DB.Save(&p).Error
}

func (pr *Produto) ProcuraPorID(id string) (*entity.Produto, error) {
	var ProdutoNovo entity.Produto

	err := pr.DB.Where("id = ?", id).First(&ProdutoNovo).Error

	return &ProdutoNovo, err
}

func (pr *Produto) Apagar(id string) error {
	p, err := pr.ProcuraPorID(id)

	if err != nil {
		return err
	}

	return pr.DB.Delete(&p).Error
}
