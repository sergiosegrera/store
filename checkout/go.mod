module github.com/sergiosegrera/store/checkout

go 1.14

replace github.com/sergiosegrera/store/cart => ../cart

require (
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/sergiosegrera/store/cart v0.0.0-00010101000000-000000000000
	github.com/stripe/stripe-go v70.15.0+incompatible
	github.com/stripe/stripe-go/v71 v71.8.0
	go.uber.org/zap v1.13.0
)
