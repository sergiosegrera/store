package service

import (
	"context"
	"log"

	//	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/cart/models"
	"github.com/sergiosegrera/store/cart/service"
)

type CheckoutService interface {
	PostCheckout(ctx context.Context, cart models.Cart) string
}

type Service struct {
	//	db *pg.DB
	cc *service.CartService
}

func NewService(c *service.CartService) *Service {
	return &Service{cc: c} //&Service{db: d, cc: c}
}

func (s *Service) PostCheckout(ctx context.Context, cart models.Cart) string {
	out := s.cc.PostCart(cart)
	log.Println(out)
	return "works"
}
