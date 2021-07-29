package dto

// TODO: Adicionar RecipeEditDTO
type RecipeDto struct {
	Name        string
	Description string
}

func NewRecipeDto() *RecipeDto {
	return &RecipeDto{}
}
