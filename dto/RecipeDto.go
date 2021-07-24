package dto

type RecipeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewRecipeDto() *RecipeDto {
	return &RecipeDto{}
}
