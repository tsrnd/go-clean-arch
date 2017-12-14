package product

import "errors"

var (
	InternalServerError = errors.New("Internal Server Error")
	NotFoundError       = errors.New("Your requested Item is not found")
	ConflictError       = errors.New("Your Item already exist")
)
