// Package that have all app controllers to resolve REST functionalities
// Every controller struct have one Repository for storage persistence purposes
package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/duducharapa/recipist-api/dto"
	"github.com/duducharapa/recipist-api/repository"
	"github.com/gorilla/mux"
)

type ProductController struct {
	Repository *repository.ProductRepository
}

func NewProductController(router *mux.Router, db *sql.DB) *ProductController {
	controller := &ProductController{Repository: repository.NewProductRepository(db)}

	router.HandleFunc("/products", controller.AllProducts).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controller.FindProduct).Methods("GET")
	router.HandleFunc("/products/{id}", controller.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/products/{id}", controller.DeleteProduct).Methods("DELETE")

	return controller
}

// Route: GET /products
// Returns:
// 	200: Product array
//	500: Failed to search any or more products
func (c *ProductController) AllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Repository.All()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
	return
}

// Route: POST /products
// Returns:
//	200: The created product
//	400: The JSON is malformed or not are JSON
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductDto
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Adicionar tratamento de erro caso o DTO falhe
	product := c.Repository.Create(&productDto)
	c.Repository.Save(product)

	json.NewEncoder(w).Encode(product)
	return
}

// TODO: Adicionar verificação de UUID
// Route: GET /products/{id}
// Returns:
//	200: The product searched by ID
//	404: Cannot find the product by this ID
func (c *ProductController) FindProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := c.Repository.Find(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
	return
}

// Route: PATCH /products/{id}
// Returns:
//	200: The updated product
//	400: The JSON is malformed or not are JSON
//	500: Failed to update the product
//  404: Product not found at last search
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var productDto dto.ProductDto
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Repository.Update(params["id"], &productDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := c.Repository.Find(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
	return
}

// Route: DELETE /products/{id}
// Returns:
//	204: The product is deleted
//	404: Cannot find the product by this ID
//	500: Failed to delete the product
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deleted, err := c.Repository.Delete(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !deleted {
		http.Error(w, errors.New("Product not deleted").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
