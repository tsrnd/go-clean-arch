package usecase

import (
	model "github.com/tsrnd/go-clean-arch/product"
	repos "github.com/tsrnd/go-clean-arch/product/repository"
)

// ProductUsecase interface
type ProductUsecase interface {
	Create(title, description string, userID int64) (*model.Product, error)
	GetByID(id int64) (*model.Product, error)
	GetByTitle(title string) (*model.Product, error)
	Update(productID int64, title, description string) (*model.Product, error)
	Delete(id int64) error
	Fetch(offset, limit int64) ([]*model.Product, error)
}

type productUsecase struct {
	repo repos.ProductRepository
}

func (uc *productUsecase) Create(title, description string, userID int64) (*model.Product, error) {
	exist, _ := uc.GetByTitle(title)
	if exist != nil {
		return nil, model.ConflictError
	}

	id, err := uc.repo.Create(title, description, userID)
	if err != nil {
		return nil, err
	}

	return uc.GetByID(id)
}

func (uc *productUsecase) GetByID(id int64) (*model.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) GetByTitle(title string) (*model.Product, error) {
	return uc.repo.GetByTitle(title)
}

func (uc *productUsecase) Update(productID int64, title, description string) (*model.Product, error) {
	err := uc.repo.Update(productID, title, description)
	if err != nil {
		return nil, err
	}
	return uc.GetByID(productID)
}

func (uc *productUsecase) Delete(id int64) error {
	existedProduct, _ := uc.GetByID(id)
	if existedProduct == nil {
		return model.NotFoundError
	}
	return uc.repo.Delete(id)
}

func (uc *productUsecase) Fetch(offset, limit int64) ([]*model.Product, error) {
	if limit == 0 {
		limit = 10
	}
	return uc.repo.Fetch(offset, limit)
}

// NewProductUsecase func
func NewProductUsecase(uc repos.ProductRepository) ProductUsecase {
	return &productUsecase{uc}
}
