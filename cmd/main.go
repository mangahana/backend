package main

import (
	"api/internal/infrastructure/repository"
	"api/pkg/configuration"
	"api/pkg/postgresql"
	"log"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.Connect(&cfg.Postgresql)
	if err != nil {
		log.Fatal(err)
	}

	_ = repository.New(db)
}
