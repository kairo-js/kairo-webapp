package main

import (
	"log"
	"net/http"

	"github.com/shizuku86/kairo-webapp/back/internal/database"
	"github.com/shizuku86/kairo-webapp/back/internal/router"
)

func main() {

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := router.NewRouter()

	log.Println("API server started :8000")

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}