package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/tsrnd/go-clean-arch/product/usecase"
	"github.com/tsrnd/go-clean-arch/services/cache"
)

// ProductController type
type ProductController struct {
  Usecase *usecase.ProductUsecase,
  Cache *cache.Cache,
}

// NewProductController func
func NewProductController(r *chi.Router, uc *usecase.ProductUsecase, c *cache.Cache) *ProductController {
	handler := &ProductController{
		Usecase: uc,
		Cache:   c,
	}
	r.POST("/users", handler.UserRegister)
	r.POST("/auth", handler.UserLogin)
	return handler
}

// Create func
func (ctrl *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	token := r.Header.Get("token")
	userIDStr, err := ctrl.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Convert user id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var cjr requests.CreateProductRequest
	err = decoder.Decode(&cjr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = ctrl.Usecase.CreateProduct(cjr.Title, cjr.Description, userID)
	if err != nil {
		log.Fatalf("Creating a product: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Product func
func (ctrl *ProductController) Product(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	productID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	product, err := ctrl.Usecase.GetProductByID(ctrl.DB, productID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
	token := r.Header.Get("token")
	userIDStr, err := ctrl.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Convert user id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if userID != product.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var ujr requests.UpdateProductRequest
		err = decoder.Decode(&ujr)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = ctrl.Usecase.UpdateProduct(ctrl.DB, product.ID, ujr.Title, ujr.Description)
		if err != nil {
			log.Fatalf("Updating a product: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "DELETE" {
		err = ctrl.Usecase.DeleteProduct(ctrl.DB, product.ID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

// Feed func
func (ctrl *ProductController) Feed(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	offset := 0
	offsetStr, ok := r.URL.Query()["offset"]
	if ok {
		offset, err = strconv.Atoi(offsetStr[0])
		if err != nil {
			offset = 0
		}
	}

	limit := 10
	limitStr, ok := r.URL.Query()["limit"]
	if ok {
		limit, err = strconv.Atoi(limitStr[0])
		if err != nil {
			limit = 1
		}
	}
	products, err := ctrl.Usecase.GetProducts(ctrl.DB, offset, limit)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
