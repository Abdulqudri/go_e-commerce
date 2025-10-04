package dtos


type CreateCategoryDTO struct {
	ID  uint `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}