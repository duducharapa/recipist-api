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

type RecipeController struct {
	Repository *repository.RecipeRepository
}

func NewRecipeController(router *mux.Router, db *sql.DB) *RecipeController {
	controller := &RecipeController{Repository: repository.NewRecipeRepository(db)}

	router.HandleFunc("/recipes", controller.AllRecipes).Methods("GET")
	router.HandleFunc("/recipes", controller.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", controller.FindRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", controller.UpdateRecipe).Methods("PATCH")
	router.HandleFunc("/recipes/{id}", controller.DeleteRecipe).Methods("DELETE")

	return controller
}

// Route: GET /recipes
// Returns:
// 	200: Recipe array
//	500: Failed to search any or more recipes
func (c *RecipeController) AllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := c.Repository.All()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recipes)
	return
}

// Route: POST /recipes
// Returns:
//	200: The created recipe
//	400: The JSON is malformed or not are JSON
func (c *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipeDto dto.RecipeDto
	err := json.NewDecoder(r.Body).Decode(&recipeDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe := c.Repository.Create(&recipeDto)
	c.Repository.Save(recipe)

	json.NewEncoder(w).Encode(recipe)
	return
}

// Route: GET /recipes/{id}
// Returns:
//	200: The recipe searched by ID
//	404: Cannot find the recipe by this ID
func (c *RecipeController) FindRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := c.Repository.Find(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(recipe)
	return
}

// Route: PATCH /recipes/{id}
// Returns:
//	200: The updated recipe
//	400: The JSON is malformed or not are JSON
//	500: Failed to update the recipe
//  404: Recipe not found at last search
func (c *RecipeController) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipeDto dto.RecipeDto
	err := json.NewDecoder(r.Body).Decode(&recipeDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Repository.Update(params["id"], &recipeDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recipe, err := c.Repository.Find(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(recipe)
	return
}

// Route: DELETE /recipes/{id}
// Returns:
//	204: The recipe is deleted
//	404: Cannot find the recipe by this ID
//	500: Failed to delete the recipe
func (c *RecipeController) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deleted, err := c.Repository.Delete(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !deleted {
		http.Error(w, errors.New("Recipe not deleted").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
