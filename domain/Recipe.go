package domain

import (
	"github.com/satori/uuid"
)

type Recipe struct {
	Id          string
	Name        string
	Description string
	Products    []Product
}

func NewRecipe(name string) *Recipe {
	return &Recipe{
		Id:   uuid.NewV4().String(),
		Name: name,
	}
}

func (r *Recipe) SetDescription(description string) {
	r.Description = description
}

/*
func (r *Recipe) Products() (string, []Product) {
	if len(r.products) == 0 {
		return "This recipe doesn't have products yet", []Product{}
	}

	return "", r.products
}

// Verificar se o produto já existe na lista
func (r *Recipe) AddProduct(product Product) {
	r.products = append(r.products, product)
}

func (r *Recipe) RemoveProduct(product Product) {
	index := -1
	for i := range r.products {
		if product.Id != r.products[i].Id {
			continue
		}

		index = i
		break
	}

	if index != -1 {
		r.products = append(r.products[:index], r.products[index+1:]...)
	}
}
*/
