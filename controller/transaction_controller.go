package controllers

import (
	"fmt"
	"log"
	model "pd-backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransaction(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order`"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	var order model.Order
	var orders []model.Order
	var orderdetail model.OrderDetail
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)

			// var orderdetails []model.OrderDetail

			detailquery := "SELECT * FROM `orderdetail` WHERE order_id='" + strconv.Itoa(order.ID) + "';"

			rows2, err2 := db.Query(detailquery)
			if err2 != nil {
				log.Println(err2)
			}

			for rows2.Next() { // loop orderdetail
				orderdetails = nil
				if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
					log.Fatal(err.Error())
				} else {
					fmt.Println(orderdetail)
					orderdetails = append(orderdetails, orderdetail)
				}
			}
			transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
			orderdetails = nil
		}
		// transaction.DataOrder = append(transaction.DataOrder, order)
		transaction.DataOrder = order
		transactions = append(transactions, transaction)
	}

	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Transactions Success"
		Response.Data = transactions
		sendHistorySuccessResponse(c, Response)
	} else {
		Response.Message = "Get Transaction Query Error"
		fmt.Print(err)
		sendHistoryErrorResponse(c, Response)
	}
}

func GetTransactionByPizza(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order` WHERE pi='0';"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

	var Response model.OrderResponse
	if err == nil {
		Response.Message = "Get Order Success"
		Response.DataOrder = orders
		sendOrderSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Order Query Error"
		fmt.Print(err)
		sendOrderErrorOResponse(c, Response)
	}
}

func GetTransactionByDate(c *gin.Context) {

}
