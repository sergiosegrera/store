package service

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product-manager/models"
)

type ProductManagerService interface {
	GetProducts(ctx context.Context) ([]*models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	PostProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, id int64) error
	// TODO: Option management
	PostOption(ctx context.Context, option models.Option) error
	// DeleteOption(ctx context.Context, id int64) error
}

type Service struct {
	db *pg.DB
}

func NewService(d *pg.DB) *Service {
	return &Service{db: d}
}

func (s *Service) GetProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := s.db.Model(&products).Select()
	if err != nil {
		return nil, err
	}

	return products, err
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

func (s *Service) PostProduct(ctx context.Context, product models.Product) error {
	err := s.db.Insert(&product)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	product := &models.Product{Id: id}
	err := s.db.Delete(product)

	return err
}

func (s *Service) PostOption(ctx context.Context, option models.Option) error {
	err := s.db.Insert(&option)
	if err != nil {
		return err
	}

	return err
}
