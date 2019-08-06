package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kevinmartinda/extra-task-week3/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/item", controllers.FindAllItem).Methods(http.MethodGet)
	router.HandleFunc("/api/item/{id}", controllers.FindItemById).Methods(http.MethodGet)
	router.HandleFunc("/api/item", controllers.CreateItem).Methods(http.MethodPost)
	router.HandleFunc("/api/item/{id}", controllers.FindItemByIdAndUpdate).Methods(http.MethodPut)
	router.HandleFunc("/api/item/{id}", controllers.FindItemByIdAndDelete).Methods(http.MethodDelete)

	router.HandleFunc("/api/kategori", controllers.FindAllCategories).Methods(http.MethodGet)
	router.HandleFunc("/api/kategori/{id}", controllers.FindCategoryById).Methods(http.MethodGet)
	router.HandleFunc("/api/kategori", controllers.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/api/kategori/{id}", controllers.FindCategoryByIdAndUpdate).Methods(http.MethodPut)
	router.HandleFunc("/api/kategori/{id}", controllers.FindCategoryByIdAndDelete).Methods(http.MethodDelete)

	log.Printf("Starting http server at port :69")
	if err := http.ListenAndServe(":69", router); err != nil {
		panic("Something went wrong..")
	}
}