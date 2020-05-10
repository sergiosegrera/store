package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	//	"github.com/go-pg/pg/v9"
	cartclient "github.com/sergiosegrera/store/cart/clients/grpc"
	"github.com/sergiosegrera/store/cart/pb"
	"github.com/stripe/stripe-go"
	stripeclient "github.com/stripe/stripe-go/client"
)

type CheckoutService interface {
	PostCheckout(ctx context.Context, cart pb.Cart) (string, error)
	PostConfirm(ctx context.Context, id string) error
}

type Service struct {
	//	db *pg.DB
	cc *cartclient.Client
	sc *stripeclient.API
}

func NewService(c *cartclient.Client, s *stripeclient.API) CheckoutService {
	return Service{cc: c, sc: s} //&Service{db: d, cc: c}
}

func (s Service) PostCheckout(ctx context.Context, cart pb.Cart) (string, error) {
	// Use Cart service to verify cart contains legal items
	cartResponse, err := s.cc.PostCart(&cart)
	if err != nil {
		return "", ErrCouldNotVerifyCart
	}

	if len(cartResponse.CartProducts) == 0 {
		return "", ErrNoProductsInCart
	}

	lineItems := []*stripe.CheckoutSessionLineItemParams{}
	for _, product := range cartResponse.CartProducts {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Name:        stripe.String(product.Name),
			Description: stripe.String(product.Description),
			Amount:      stripe.Int64(product.Price),
			Currency:    stripe.String(string(stripe.CurrencyCAD)),
			Quantity:    stripe.Int64(product.Count),
		})
	}

	// TODO: Get allowed countries service
	// TODO: Change redirect urls to match frontend
	paymentIntentDataParams := &stripe.CheckoutSessionPaymentIntentDataParams{
		CaptureMethod: stripe.String("manual"),
	}

	// Add cart to context/metadata
	cartJson, err := json.Marshal(cartResponse)
	if err != nil {
		return "", err
	}

	paymentIntentDataParams.AddMetadata("cart", string(cartJson))

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems:         lineItems,
		PaymentIntentData: paymentIntentDataParams,
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: []*string{
				stripe.String("US"),
				stripe.String("CA"),
			},
		},
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
	}

	session, err := s.sc.CheckoutSessions.New(params)
	if err != nil {
		return "", ErrCreatingToken
	}

	return session.ID, err
}

func (s Service) PostConfirm(ctx context.Context, id string) error {
	paymentIntent, err := s.sc.PaymentIntents.Get(id, nil)
	if err != nil {
		return err
	}

	log.Println(paymentIntent.Metadata)
	return err
}

var (
	ErrCouldNotVerifyCart = errors.New("Could not verify cart")
	ErrNoProductsInCart   = errors.New("No products in cart")
	ErrCreatingToken      = errors.New("Error creating token")
)
