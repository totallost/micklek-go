package main

import (
	"controllers"
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

var err error

//DB is a global database variable
var DB *sql.DB

func main() {
	DB, err = sql.Open("sqlite3", "./Micklek.db")
	controllers.DB = DB
	defer DB.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/items/all", controllers.GetAllItems).Methods("GET")
	router.HandleFunc("/api/items", controllers.GetActiveItems).Methods("GET")
	router.HandleFunc("/api/items/{id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/api/items/{id}", controllers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/api/items/{id}", controllers.UpdateItem).Methods("POST")
	router.HandleFunc("/api/items/add", controllers.AddNewItem).Methods("PUT")

	router.HandleFunc("/api/management/get-statuses", controllers.GetStatuses).Methods("GET")
	router.HandleFunc("/api/management/get-order-headers", controllers.GetOrderdHeaders).Methods("GET")
	router.HandleFunc("/api/management/get-order-header/{id}", controllers.GetOrderdHeader).Methods("GET")
	router.HandleFunc("/api/management/get-order-lines/{id}", controllers.GetOrderLines).Methods("GET")
	router.HandleFunc("/api/details/sendOrder", controllers.CreateNewOrder).Methods("PUT")
	router.HandleFunc("/api/management/update-order", controllers.UpdateOrder).Methods("POST")

	http.ListenAndServe(":8080", router)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
