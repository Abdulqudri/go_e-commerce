package dtos


type CreateProductDTO struct {
	ID          uint      `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Category    string    `json:"category" binding:"required"`
}

type UpdateProductDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}