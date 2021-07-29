package domain

type Ingredient struct {
	RecipeId  string
	ProductId string
	Quantity  int32
}

func NewIngredient() *Ingredient {
	return &Ingredient{}
}
