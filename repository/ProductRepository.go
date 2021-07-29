package repository

import (
	"database/sql"

	"github.com/duducharapa/recipist-api/domain"
	"github.com/duducharapa/recipist-api/dto"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) All() ([]*domain.Product, error) {
	products := make([]*domain.Product, 0)
	rows, err := r.Db.Query("SELECT id,name,description,quantity FROM products")
	defer rows.Close()

	if err != nil {
		return products, err
	}

	for rows.Next() {
		product := new(domain.Product)

		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Quantity); err != nil {
			return make([]*domain.Product, 0), err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) Create(product *dto.ProductDto) *domain.Product {
	newProduct := domain.NewProduct(product.Name, product.Quantity)
	if product.Description != "" {
		newProduct.SetDescription(product.Description)
	}

	return newProduct
}

func (r *ProductRepository) Save(product *domain.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products(id,name,description,quantity) VALUES($1,$2,$3,$4)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		product.Id,
		product.Name,
		product.Description,
		product.Quantity,
	)

	if err != nil {
		return err
	}

	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Find(id string) (domain.Product, error) {
	var p domain.Product
	stmt, err := r.Db.Prepare("SELECT id,name,description,quantity FROM products WHERE id=$1 LIMIT 1")

	if err != nil {
		return p, err
	}

	if err = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Description, &p.Quantity); err != nil {
		return p, err
	}

	return p, nil
}

func (r *ProductRepository) Update(id string, productDto *dto.ProductDto) (domain.Product, error) {
	_, err := r.Db.Exec("UPDATE products SET name=$1,description=$2,quantity=$3 WHERE id=$4",
		productDto.Name, productDto.Description, productDto.Quantity, id,
	)

	if err != nil {
		return domain.Product{}, err
	}

	product, err := r.Find(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id string) error {
	_, err := r.Db.Exec("DELETE FROM products WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}
