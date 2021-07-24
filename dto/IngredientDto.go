package dto

type IngredientDto struct {
	RecipeId  string
	ProductId string
	Quantity  int32
}

func NewIngredientDto() *IngredientDto {
	return &IngredientDto{}
}
