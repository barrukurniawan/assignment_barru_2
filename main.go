package main

import (
	"assignment_2/db"
	"assignment_2/server"
	"assignment_2/server/controllers"
)

// Assignment 2
// Barru Kurniawan

func main() {
	db := db.ConnectGorm()
	orderController := controllers.NewOrderController(db)
	router := server.NewRouter(orderController)
	router.Start(":4000")
}
