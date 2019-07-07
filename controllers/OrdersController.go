package controllers

import (
	"dtos"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//GetStatuses get you the Statuses
func GetStatuses(w http.ResponseWriter, r *http.Request) {
	status := models.Status{}
	statuses := []models.Status{}
	rows, err := DB.Query("select * from statuses")
	check(err)
	for rows.Next() {
		rows.Scan(&status.Id, &status.Description)
		statuses = append(statuses, status)
	}
	answer, _ := json.Marshal(statuses)
	fmt.Fprint(w, string(answer))
}

//GetOrderdHeaders Gets all Order Headers
func GetOrderdHeaders(w http.ResponseWriter, r *http.Request) {
	oh := models.OrderHeader{}
	ohs := []models.OrderHeader{}

	rows, err := DB.Query("select * from orderheaders")
	check(err)

	for rows.Next() {
		rows.Scan(&oh.Id, &oh.NumberOfItems, &oh.TotalPrice, &oh.ClientFirstName, &oh.ClientSureName, &oh.ClientEmail, &oh.ClientCell, &oh.ClientRemarks, &oh.DateCreation, &oh.DateTarget, &oh.StatusId, &oh.Address)
		ohs = append(ohs, oh)
	}
	answer, _ := json.Marshal(ohs)
	fmt.Fprint(w, string(answer))
}

//GetOrderdHeader get one header
func GetOrderdHeader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	oh := models.OrderHeader{}

	_ = DB.QueryRow("select * from orderheaders where id=?", params["id"]).Scan(
		&oh.Id,
		&oh.NumberOfItems,
		&oh.TotalPrice,
		&oh.ClientFirstName,
		&oh.ClientSureName,
		&oh.ClientEmail,
		&oh.ClientCell,
		&oh.ClientRemarks,
		&oh.DateCreation,
		&oh.DateTarget,
		&oh.StatusId,
		&oh.Address,
	)
	answer, err := json.Marshal(oh)
	check(err)
	fmt.Fprint(w, string(answer))
}

//GetOrderLines gets order lines
func GetOrderLines(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ol := models.OrderLine{}
	ols := []models.OrderLine{}

	rows, err := DB.Query("select * from orderlines where orderheaderid=?", params["id"])
	check(err)
	for rows.Next() {
		rows.Scan(
			&ol.OrderHeaderId,
			&ol.LineNumber,
			&ol.ItemId,
			&ol.Amount,
			&ol.LinePrice,
		)
		ols = append(ols, ol)
	}

	answer, err := json.Marshal(ols)
	check(err)
	fmt.Fprint(w, string(answer))
}

//CreateNewOrder creates new order in the db
func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	var orderDetailsDto dtos.OrderDetailsDto
	err := json.NewDecoder(r.Body).Decode(&orderDetailsDto)
	check(err)
	if orderDetailsDto.ClienDetails.SureName == "" {
		panic(http.StatusBadRequest)
	}
	var numberOfItems int
	var totalPrice float64
	for _, v := range orderDetailsDto.OrderDetails {
		numberOfItems = numberOfItems + v.Amount
		totalPrice = totalPrice + float64(v.Amount)*v.LinePrice
	}
	rows, err := DB.Prepare("insert into OrderHeaders (numberOfItems, totalPrice, ClientFirstName,ClientSureName,ClientEmail,ClientCell,ClientRemarks,DateCreation,DateTarget,StatusId) values(?,?,?,?,?,?,?,?,?,?)")
	check(err)
	re, err := rows.Exec(numberOfItems,
		totalPrice,
		orderDetailsDto.ClienDetails.FirstName,
		orderDetailsDto.ClienDetails.SureName,
		orderDetailsDto.ClienDetails.Email,
		orderDetailsDto.ClienDetails.MobileNumber,
		orderDetailsDto.ClienDetails.Notes,
		time.Now(),
		orderDetailsDto.ClienDetails.DateReady,
		1)
	check(err)
	LastOrderID, err := re.LastInsertId()
	check(err)
	for _, v := range orderDetailsDto.OrderDetails {
		_, err := DB.Exec("insert into OrderLines values (?,?,?,?,?)",
			LastOrderID,
			v.LineNumber,
			v.ItemID,
			v.Amount,
			v.LinePrice)
		check(err)
	}
}

//UpdateOrder updates order from managment
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var orderDetailsDto dtos.OrderDetailsDto
	err := json.NewDecoder(r.Body).Decode(&orderDetailsDto)
	check(err)
	if orderDetailsDto.ClienDetails.FirstName == "" {
		panic(http.StatusBadRequest)
	}

	//delete old order lines
	_, err = DB.Exec("delete from OrderLines where orderHeaderId=?", orderDetailsDto.ClienDetails.ID)
	check(err)

	var numberOfItems int
	var totalPrice float64
	for _, v := range orderDetailsDto.OrderDetails {
		numberOfItems = numberOfItems + v.Amount
		totalPrice = totalPrice + float64(v.Amount)*v.LinePrice
	}
	rows, err := DB.Prepare("update OrderHeaders set(numberOfItems, totalPrice, ClientFirstName,ClientSureName,ClientEmail,ClientCell,ClientRemarks,DateCreation,DateTarget,StatusId) values(?,?,?,?,?,?,?,?,?,?) where Id=?")
	check(err)
	re, err := rows.Exec(numberOfItems,
		totalPrice,
		orderDetailsDto.ClienDetails.FirstName,
		orderDetailsDto.ClienDetails.SureName,
		orderDetailsDto.ClienDetails.Email,
		orderDetailsDto.ClienDetails.MobileNumber,
		orderDetailsDto.ClienDetails.Notes,
		time.Now(),
		orderDetailsDto.ClienDetails.DateReady,
		1,
		orderDetailsDto.ClienDetails.ID)
	check(err)
	LastOrderID, err := re.LastInsertId()
	check(err)
	for _, v := range orderDetailsDto.OrderDetails {
		_, err := DB.Exec("insert into OrderLines values (?,?,?,?,?)",
			LastOrderID,
			v.LineNumber,
			v.ItemID,
			v.Amount,
			v.LinePrice)
		check(err)
	}
}
