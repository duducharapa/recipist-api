package domain

import "github.com/satori/uuid"

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int32  `json:"quantity"`
}

func NewProduct(name string, quantity int32) *Product {
	return &Product{
		Id:       uuid.NewV4().String(),
		Name:     name,
		Quantity: quantity,
	}
}

func (p *Product) SetDescription(description string) {
	p.Description = description
}

func (p *Product) Consume(quantity int32) {
	if quantity <= p.Quantity {
		p.Quantity -= quantity
	}
}
