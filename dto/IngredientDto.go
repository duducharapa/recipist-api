package dto

type IngredientDto struct {
	RecipeId  string
	ProductId string
	Quantity  int32
}

// TODO: Adicionar IngredientEditDTO
func NewIngredientDto() *IngredientDto {
	return &IngredientDto{}
}
