package main

import (
	"API-GOKOMODO/controllers"
	"API-GOKOMODO/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

/*
	Main function.

*
*/
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Initialize Database
	database.GetDatabase()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// // Register Routes
	RegistertRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", os.Getenv("APP_PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router))
}

func RegistertRoutes(router *mux.Router) {
	// Auth
	router.HandleFunc("/api/auth/buyer", controllers.LoginBuyer).Methods("POST")
	router.HandleFunc("/api/auth/seller", controllers.LoginSeller).Methods("POST")
	// Product
	router.HandleFunc("/api/products", controllers.IsAuthorized(controllers.GetProducts)).Methods("GET")
	router.HandleFunc("/api/products", controllers.IsAuthorized(controllers.CreateProduct)).Methods("POST")
	// Order
	router.HandleFunc("/api/orders", controllers.IsAuthorized(controllers.GetOrders)).Methods("GET")
	router.HandleFunc("/api/orders", controllers.IsAuthorized(controllers.CreateOrder)).Methods("POST")
	router.HandleFunc("/api/orders/update/status/{id}", controllers.IsAuthorized(controllers.UpdateStatusOrder)).Methods("PUT")
}
