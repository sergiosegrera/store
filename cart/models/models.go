package models

type Cart struct {
	CartProducts []*CartProduct `json:"cartProducts"`
}

type CartProduct struct {
	Id       int64 `json:"id"`
	OptionId int64 `json:"optionId"`
	Count    int64 `json:"count"`
	Price    int64 `json:"price"`
}
