package entity

import "github.com/google/uuid"

type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category {
		ID: uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID	string  	`json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CategoryID string `json:"category_id"`
	 // Poderia ser um Category 
	ImageURL string  `json:"image_url"`
}

 //se nao coloco o tipo ele coloca como string 
func NewProduct (name, description string, categoryID, imageURL string, price float64) *Product {
	return &Product { 
		ID: uuid.New().String(),
		Name: name,
		Description: description, 
		Price: price,
		CategoryID: categoryID,
		ImageURL: imageURL,
	}

}