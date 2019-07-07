package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

//DB global database variable
var DB *sql.DB

//GetAllItems gets all items
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	items := []models.Item{}

	rows, err := DB.Query("select * from items")
	check(err)

	for rows.Next() {
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&item.PhotoUrl,
			&item.Description,
			&item.IsActive,
			&item.PhotoPublicName,
		)
		items = append(items, item)
	}
	answer, _ := json.Marshal(items)
	fmt.Fprint(w, string(answer))

	defer rows.Close()
}

//GetActiveItems gets all active items with isActive=1
func GetActiveItems(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	items := []models.Item{}

	rows, err := DB.Query("select * from items where isActive=1")
	check(err)

	for rows.Next() {
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&item.PhotoUrl,
			&item.Description,
			&item.IsActive,
			&item.PhotoPublicName,
		)
		items = append(items, item)
	}
	answer, _ := json.Marshal(items)
	fmt.Fprint(w, string(answer))

	defer rows.Close()
}

//GetItem get you one item
func GetItem(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	params := mux.Vars(r)
	_ = DB.QueryRow("select * from items where Id= ?", params["id"]).Scan(
		&item.Id,
		&item.Name,
		&item.Price,
		&item.PhotoUrl,
		&item.Description,
		&item.IsActive,
		&item.PhotoPublicName,
	)

	answer, _ := json.Marshal(item)
	fmt.Fprint(w, string(answer))
}

//DeleteItem deltes one item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rows, err := DB.Prepare("delete from items where id=?")
	check(err)
	re, er := rows.Exec(params["id"])
	check(er)
	n, err := re.RowsAffected()
	check(err)
	defer rows.Close()

	fmt.Fprint(w, "item deleted", n)

}

//UpdateItem updates one item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item := models.Item{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	check(err)

	rows, err := DB.Prepare("update items set name=?, price=?, PhotoUrl=?, Description=?, IsActive=?,PhotoPublicName=? where id=?")
	check(err)
	re, err := rows.Exec(item.Name, item.Price, item.PhotoUrl, item.Description, item.IsActive, item.PhotoPublicName, params["id"])
	check(err)
	n, err := re.RowsAffected()
	check(err)

	fmt.Fprint(w, n, " rows affected")
}

//AddNewItem adds a new item
func AddNewItem(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&item)
	check(err)

	rows, err := DB.Prepare("insert into items (Name,Price,PhotoUrl,Description,IsActive,PhotoPublicName) values (?,?,?,?,?,?)")
	check(err)
	re, err := rows.Exec(item.Name, item.Price, item.PhotoUrl, item.Description, item.IsActive, item.PhotoPublicName)
	check(err)
	n, err := re.RowsAffected()
	fmt.Fprint(w, n, " Item added")
	fmt.Fprintln(w, item.Id)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
