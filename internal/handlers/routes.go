package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Chintukr2004/student-api/internal/middleware"
	"github.com/Chintukr2004/student-api/internal/repository"
	"github.com/Chintukr2004/student-api/internal/service"
	"github.com/go-chi/chi/v5"
)

func Routes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	//global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	UserHandler := NewUserHandler(userService)

	r.Post("/users", UserHandler.Register)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	return r
}
