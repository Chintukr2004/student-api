package main

import (
	"log"

	"github.com/Chintukr2004/student-api/internal/app"
	"github.com/Chintukr2004/student-api/internal/config"
)

func main() {
	cfg := config.Load()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
