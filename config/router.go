package config

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	productRepo "github.com/tsrnd/go-clean-arch/product/repository"
	productCase "github.com/tsrnd/go-clean-arch/product/usecase"
	"github.com/tsrnd/go-clean-arch/services/cache"
	userRepo "github.com/tsrnd/go-clean-arch/user/repository"
	userCase "github.com/tsrnd/go-clean-arch/user/usecase"
)

// Router func
func Router(db *sql.DB, c cache.Cache) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addUserRoutes(r, db, c)
	addProductRoutes(r, db, c)
}

func addUserRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

func addProductRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := productRepo.NewUserRepository(db)
	uc := productCase.NewUserUsecase(repo)
	productDeliver.NewProductController(r, uc, c)
}
