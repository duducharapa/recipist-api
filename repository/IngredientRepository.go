package repository

import (
	"database/sql"
	"errors"

	"github.com/duducharapa/recipist-api/domain"
)

type IngredientRepository struct {
	Db *sql.DB
}

func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{Db: db}
}

func (r *IngredientRepository) Find(productId string, recipeId string) (domain.Ingredient, error) {
	var i domain.Ingredient
	stmt, err := r.Db.Prepare("SELECT product_id, recipe_id, quantity FROM ingredients WHERE product_id=$1 AND recipe_id=$2")

	if err != nil {
		return i, err
	}

	if err = stmt.QueryRow(productId, recipeId).Scan(&i.ProductId, &i.RecipeId, &i.Quantity); err != nil {
		return i, errors.New("Ingredient not included on this recipe")
	}

	return i, nil
}

/*
func (r *IngredientRepository) Add(ingredient *domain.Ingredient) error {
	stmt, err := r.Db.Prepare("INSERT INTO ingredients(product_id, recipe_id, quantity) VALUES($1,$2,$3)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(ingredient.ProductId, ingredient.RecipeId, ingredient.Quantity)

	if err != nil {
		return err
	}
}
*/
