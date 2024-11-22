package main

import (
	"log"
	"net/http"
	"perpustakaan/config"
	"perpustakaan/internal/api"
	"perpustakaan/pkg/database"
)

func main() {
	config.Load()
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	r := api.NewRouter()
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
