// TODO: Update Host on Swagger
// TODO: Ler sobre a licença MIT
// TODO: Adicionar funcionalidade de APIKEY na aplicação e na documentação

//	this application is a REST API made in Go for the Recipist application
//
//	Terms Of Service:
//
//	there are no TOS at this moment, use at your own risk we take no responsability
//
//		Schemes: http
//		Host: localhost
//		BasePath: /
//		Version: 1.0
//		License: MIT http://opensource.org/licenses/MIT
//		Contact: Eduardo Charapa<eduardocharapa@gmail.com> https://github.com/duducharapa
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := mux.NewRouter()
	db := database.SetupDB()

	controllers.NewProductController(router, db)
	controllers.NewRecipeController(router, db)

	log.Fatal(http.ListenAndServe(os.Getenv("appport"), router))
}
