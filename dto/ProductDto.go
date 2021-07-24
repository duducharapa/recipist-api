package dto

type ProductDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int32  `json:"quantity"`
}

func NewProductDto() *ProductDto {
	return &ProductDto{}
}
