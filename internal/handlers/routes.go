package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chintukr2004/student-api/internal/config"
	"github.com/Chintukr2004/student-api/internal/middleware"
	"github.com/Chintukr2004/student-api/internal/repository"
	"github.com/Chintukr2004/student-api/internal/service"
	"github.com/go-chi/chi/v5"
)

func Routes(db *sql.DB, cfg config.Config) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService, cfg.JWT.Secret)

	r.Post("/login", userHandler.Login)

	r.Post("/users", userHandler.Register)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(cfg.JWT.Secret))

		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			userID := r.Context().Value(middleware.UserIDKey)
			role := r.Context().Value(middleware.RoleKey)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"user_id": userID,
				"role":    role,
			})
		})
	})

	return r
}
