package psql

import (
	"database/sql"

	model "github.com/tsrnd/go-clean-arch/product"
	repo "github.com/tsrnd/go-clean-arch/product/repository"
)

type productRepository struct {
	DB *sql.DB
}

func (m *productRepository) Create(title, description string, userID int64) (int64, error) {
	const query = `
		insert into products (
			title,
			description,
			user_id
		) values (
			$1,
			$2,
			$3
		) returning id
	`
	var id int64
	err := m.DB.QueryRow(query, title, description, userID).Scan(&id)
	return id, err
}

func (m *productRepository) GetByID(id int64) (*model.Product, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			products
		where
			id = $1
	`
	var product model.Product
	err := m.DB.QueryRow(query, id).Scan(&product.ID, &product.Title, &product.Description, &product.UserID)
	return &product, err
}

func (m *productRepository) GetByTitle(title string) (*model.Product, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			products
		where
			title = $1
	`
	var product model.Product
	err := m.DB.QueryRow(query, title).Scan(&product.ID, &product.Title, &product.Description, &product.UserID)
	return &product, err
}

func (m *productRepository) Update(productID int64, title, description string) error {
	const query = `
    update products set
			title = $1,
			description = $2
    where
      id = $3
	`
	_, err := m.DB.Exec(query, title, description, productID)
	return err
}

func (m *productRepository) Delete(id int64) error {
	const query = `delete from products where id = $1`
	_, err := m.DB.Exec(query, id)
	return err
}

func (m *productRepository) Fetch(offset, limit int64) ([]*model.Product, error) {
	const query = `
		select
			id,
			title,
			description,
			user_id
		from
			products
		limit $1 offset $2
	`
	products := make([]*model.Product, 0)

	rows, err := m.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.Title, &product.Description, &product.UserID)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, err
}

// NewProductRepository func
func NewProductRepository(DB *sql.DB) repo.ProductRepository {
	return &productRepository{DB}
}
