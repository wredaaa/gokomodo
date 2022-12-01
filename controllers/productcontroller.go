package controllers

import (
	"API-GOKOMODO/database"
	"API-GOKOMODO/entities"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	w.Header().Set("Content-Type", "application/json")

	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)

	connection.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var products []entities.Product
	connection.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
