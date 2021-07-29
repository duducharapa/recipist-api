package dto

// TODO: Adicionar ProductEditDTO
type ProductDto struct {
	Name        string
	Description string
	Quantity    int32
}

func NewProductDto() *ProductDto {
	return &ProductDto{}
}
