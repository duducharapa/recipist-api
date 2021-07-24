package main

import (
	"log"
	"net/http"
	"os"

	"github.com/duducharapa/recipist-api/controllers"
	"github.com/duducharapa/recipist-api/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// That function resolve the .ENV variables to use on application
// Like appport or host
//
// See .env or .env.example file for more details
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := mux.NewRouter()
	db := database.SetupDB()

	// Register the controllers from package CONTROLLERS
	controllers.NewProductController(router, db)
	controllers.NewRecipeController(router, db)

	log.Fatal(http.ListenAndServe(os.Getenv("appport"), router))
}
