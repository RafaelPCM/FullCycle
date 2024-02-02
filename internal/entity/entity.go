package entity

import "github.com/google/uuid"

type Category struct {
	ID string
	Name string
}

func NewCategory(name string) *Category {
	return &Category {
		ID: uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID	string
	Name string
	Description string
	Price float64
	CategoryID string // Poderia ser um Category 
	ImageURL string
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