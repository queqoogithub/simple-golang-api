package main

import (
	"log"
	"net/http"
	"simple-crud-api/handlers"
	"simple-crud-api/services"
)

func main() {
	itemService := services.NewMockItemService()

	// 🍩 Uncomment the following line to use real service
	// itemService := services.NewItemService()

	itemHandler := handlers.NewItemHandler(itemService)

	http.HandleFunc("/items", itemHandler.HandleItems)
	http.HandleFunc("/items/", itemHandler.HandleItem)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
