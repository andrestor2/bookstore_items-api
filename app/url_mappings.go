package app

import (
	"github.com/andrestor2/bookstore_items-api/controller"
	"net/http"
)

func mapUrls() {

	router.HandleFunc("/ping", controller.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controller.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controller.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controller.ItemsController.Search).Methods(http.MethodPost)
}
