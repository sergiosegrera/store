package service

import (
	"context"

	"github.com/sergiosegrera/store/product/models"

	"github.com/go-pg/pg/v9"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]*models.Thumbnail, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
}

type Service struct {
	db *pg.DB
}

func NewService(d *pg.DB) *Service {
	return &Service{db: d}
}

func (s *Service) GetProducts(ctx context.Context) ([]*models.Thumbnail, error) {
	var products []models.Product
	err := s.db.Model(&products).Where("public = true").Select()
	if err != nil {
		return nil, err
	}

	var thumbnails []*models.Thumbnail
	for _, product := range products {
		thumbnails = append(thumbnails, &models.Thumbnail{
			Id:        product.Id,
			Name:      product.Name,
			Thumbnail: product.Thumbnail,
			Price:     product.Price,
		})
	}

	return thumbnails, err
}

func (s *Service) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	product := &models.Product{Id: id}
	err := s.db.Select(product)
	if err != nil {
		return nil, err
	}

	var options []*models.Option
	err = s.db.Model(&options).Where("product_id = ?", product.Id).Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}

	product.Options = options

	return product, err
}
