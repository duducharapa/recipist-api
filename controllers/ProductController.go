package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/duducharapa/recipist-api/dto"
	"github.com/duducharapa/recipist-api/repository"
	"github.com/duducharapa/recipist-api/utils"
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

//	swagger:route GET /products products allProducts
//
//	Lists all products in the API
//
//		Produces:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//			200: Array of products
//			500: Internal error
func (c *ProductController) AllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Repository.All()

	if err != nil {
		utils.SendError(w, "Failed to search for products", http.StatusInternalServerError)
		return
	}

	utils.Send(w, products, http.StatusAccepted)
	return
}

//	swagger:route POST /products products createProduct
//
//	Create a new product
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//			201: Created product
//			400: Malformed JSON
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.ProductDto
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		utils.SendError(w, "The body needs be a JSON content!", http.StatusBadRequest)
		return
	}

	// TODO: Adicionar tratamento de erro caso o DTO falhe
	product := c.Repository.Create(&productDto)
	c.Repository.Save(product)

	utils.Send(w, product, http.StatusCreated)
	return
}

// TODO: Adicionar verificação de UUID
//	swagger:route GET /products/{id} products findProduct
//
//	Find one product of a given ID
//
//		Produces:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//			200: Product
//			404: Product not found
//			500: Internal error
func (c *ProductController) FindProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := c.Repository.Find(params["id"])

	if err != nil {
		utils.SendError(w, "Failed to search for a product", http.StatusInternalServerError)
		return
	}

	if product.IsNull() {
		utils.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.Send(w, product, http.StatusOK)
	return
}

//	swagger:route PATCH /products/{id} products updateProduct
//
//	Update a product of a given ID
//	If product not exists, throw a error before try to update
//
//		Produces:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//			200: Product updated
//			400: Malformed JSON
//			404: Product not found
//			500: Internal error
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var productDto dto.ProductDto
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		utils.SendError(w, "Body needs to be a JSON", http.StatusBadRequest)
		return
	}

	product, err := c.Repository.Find(params["id"])

	if err != nil {
		utils.SendError(w, "Failed to search for a product", http.StatusNotFound)
		return
	}

	product, err = c.Repository.Update(params["id"], &productDto)

	if err != nil {
		utils.SendError(w, "Failed to update a product", http.StatusInternalServerError)
		return
	}

	utils.Send(w, product, http.StatusOK)
	return
}

//	swagger:route DELETE /products/{id} products deleteProduct
//
//	Delete one product of a given ID
//
//		Produces:
//		- application/json
//
//		Schemes: http
//
//		Responses:
//			204: Deleted
//			404: Product not found
//			500: Internal error to find or delete
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product, err := c.Repository.Find(params["id"])

	if err != nil {
		utils.SendError(w, "Failed to find a product to delete", http.StatusInternalServerError)
		return
	}

	if product.IsNull() {
		utils.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	err = c.Repository.Delete(params["id"])

	if err != nil {
		utils.SendError(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	utils.Send(w, nil, http.StatusNoContent)
	return
}
