package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Chintukr2004/student-api/internal/config"
	"github.com/Chintukr2004/student-api/internal/database"
	"github.com/Chintukr2004/student-api/internal/handlers"
)

type App struct {
	Config config.Config
	DB     *sql.DB
}

func New(cfg config.Config) (*App, error) {
	db, err := database.Open(cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	return &App{
		Config: cfg,
		DB:     db,
	}, nil
}

func (a *App) Run() error {
	router := handlers.Routes(a.DB, a.Config)

	log.Printf("starting server on :%s", a.Config.Port)
	return http.ListenAndServe(":"+a.Config.Port, router)
}
