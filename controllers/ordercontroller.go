package controllers

import (
	"API-GOKOMODO/database"
	"API-GOKOMODO/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	w.Header().Set("Content-Type", "application/json")

	var order entities.Order
	json.NewDecoder(r.Body).Decode(&order)

	connection.Create(&order)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var orders []entities.Order
	connection.Find(&orders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	productId := mux.Vars(r)["id"]
	var order entities.Order
	connection.First(&order, productId)
	if order.ID == 0 {
		json.NewEncoder(w).Encode("Order Not Found!")
		return
	}
	json.NewDecoder(r.Body).Decode(&order)

	connection.Save(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
