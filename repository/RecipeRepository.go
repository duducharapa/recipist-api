package repository

import (
	"database/sql"
	"errors"

	"github.com/duducharapa/recipist-api/domain"
	"github.com/duducharapa/recipist-api/dto"
)

type RecipeRepository struct {
	Db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{Db: db}
}

func (r *RecipeRepository) All() ([]*domain.Recipe, error) {
	emptyRecipes := make([]*domain.Recipe, 0)
	recipes := make([]*domain.Recipe, 0)
	rows, err := r.Db.Query("SELECT id,name,description FROM recipes")

	if err != nil {
		return recipes, err
	}
	defer rows.Close()

	for rows.Next() {
		recipe := new(domain.Recipe)
		if err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Description); err != nil {
			return emptyRecipes, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepository) Create(recipe *dto.RecipeDto) *domain.Recipe {
	newRecipe := domain.NewRecipe(recipe.Name)
	if recipe.Description != "" {
		newRecipe.SetDescription(recipe.Description)
	}

	return newRecipe
}

func (r *RecipeRepository) Save(product *domain.Recipe) error {
	stmt, err := r.Db.Prepare("INSERT INTO recipes(id,name,description) VALUES($1,$2,$3)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		product.Id,
		product.Name,
		product.Description,
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

func (r *RecipeRepository) Find(id string) (domain.Recipe, error) {
	var rc domain.Recipe
	stmt, err := r.Db.Prepare("SELECT id,name,description FROM recipes WHERE id=$1 LIMIT 1")

	if err != nil {
		return rc, err
	}

	if err = stmt.QueryRow(id).Scan(&rc.Id, &rc.Name, &rc.Description); err != nil {
		return rc, errors.New("Recipe does not exist")
	}

	return rc, nil
}

func (r *RecipeRepository) Update(id string, recipeDto *dto.RecipeDto) error {
	_, err := r.Db.Exec("UPDATE recipes SET name=$1,description=$2 WHERE id=$3",
		recipeDto.Name, recipeDto.Description, id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RecipeRepository) Delete(id string) (bool, error) {
	res, err := r.Db.Exec("DELETE FROM recipes WHERE id=$1", id)

	if err != nil {
		return false, errors.New("Failed to delete product")
	}

	count, err := res.RowsAffected()

	if err != nil {
		return false, errors.New("Failed to delete product")
	}

	if count > 0 {
		return true, nil
	}

	return false, errors.New("Product not found to delete")
}
