package service

import (
	"context"

	"github.com/go-pg/pg/v9"

	"github.com/sergiosegrera/store/cart/config"
	"github.com/sergiosegrera/store/cart/pb"
	productmodels "github.com/sergiosegrera/store/product/models"
)

type CartService interface {
	PostCart(ctx context.Context, cart pb.Cart) (pb.Cart, error)
}

type Service struct {
	db     *pg.DB
	Config *config.Config
}

func NewService(d *pg.DB, c *config.Config) CartService {
	return Service{db: d, Config: c}
}

func (s Service) PostCart(ctx context.Context, cart pb.Cart) (pb.Cart, error) {
	// Check every product in cart, check if the stock is available and if the product exists.
	// Also Calculate price for every product.
	var outputCart pb.Cart
	for _, cartProduct := range cart.CartProducts {
		if cartProduct.Count > 0 {
			product := &productmodels.Product{}
			err := s.db.Model(product).Where("id = ? AND public = true", cartProduct.Id).Select()

			// TODO: Check if error is because no row found or db error
			// Cause if error exists this will panic
			// If no error, id exists, proceed
			if err == nil {
				option := &productmodels.Option{}
				err := s.db.Model(option).Where(
					"id = ? AND product_id = ?",
					cartProduct.OptionId,
					cartProduct.Id,
				).Select()

				// If no error, id exists, proceed
				if err == nil {
					// Check if there is enough stock
					if option.Stock >= cartProduct.Count {
						outputCart.CartProducts = append(outputCart.CartProducts, &pb.CartProduct{
							Id:          cartProduct.Id,
							Name:        cartProduct.Name,
							Description: cartProduct.Description,
							Thumbnail:   cartProduct.Thumbnail,
							OptionId:    cartProduct.OptionId,
							OptionName:  option.Name,
							Count:       cartProduct.Count,
							Price:       product.Price,
						})
					} else {
						outputCart.CartProducts = append(outputCart.CartProducts, &pb.CartProduct{
							Id:          cartProduct.Id,
							Name:        cartProduct.Name,
							Description: cartProduct.Description,
							Thumbnail:   cartProduct.Thumbnail,
							OptionId:    cartProduct.OptionId,
							OptionName:  option.Name,
							Count:       option.Stock,
							Price:       product.Price,
						})
					}
				}
			}
		}
	}
	return outputCart, nil
}
