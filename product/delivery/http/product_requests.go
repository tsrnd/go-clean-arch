package http

type CreateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
