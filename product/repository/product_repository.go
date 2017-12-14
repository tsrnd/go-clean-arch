package repository

import (
	model "github.com/tsrnd/go-clean-arch/product"
)

// ProductRepository interface
type ProductRepository interface {
	Create(title, description string, userID int64) (int64, error)
	Update(productID int64, title, description string) error
	Delete(id int64) error
	GetByID(id int64) (*model.Product, error)
	GetByTitle(title string) (*model.Product, error)
	Fetch(offset, limit int64) ([]*model.Product, error)
}
