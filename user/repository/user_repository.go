package repository

import (
	"database/sql"

	"github.com/tsrnd/go-clean-arch/services/crypto"
	model "github.com/tsrnd/go-clean-arch/user"
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetPrivateDetailsByEmail(email string) (*model.PrivateUserDetails, error)
	Create(email, name, password string) (int, error)
}

type userRepository struct {
	DB *sql.DB
}

func (m *userRepository) GetByID(id int64) (*model.User, error) {
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      id = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userRepository) GetByEmail(email string) (*model.User, error) {
	const query = `
    select
      id,
      email,
      name
    from
      users
    where
      email = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

func (m *userRepository) GetPrivateUserDetailsByEmail(email string) (*model.PrivateUserDetails, error) {
	const query = `
    select
      id,
      password,
      salt
    from
      users
    where
      email = $1
  `
	var u model.PrivateUserDetails
	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Password, &u.Salt)
	return &u, err
}

func (m *userRepository) Create(email, name, password string) (int64, error) {
	const query = `
    insert into users (
      email,
      name,
      password,
      salt
    ) values (
      $1,
      $2,
      $3,
      $4
    )
    returning id
  `
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int64
	err := m.DB.QueryRow(query, email, name, hashedPassword, salt).Scan(&id)
	return id, err
}
