package controllers

import (
	"fmt"
	"log"
	model "pd-backend/model"
	"strconv"
	"strings"

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
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)

			// var orderdetails []model.OrderDetail
			// detailquery := "SELECT * FROM `orderdetail` WHERE order_id='" + strconv.Itoa(order.ID) + "';"
			detailquery := "SELECT * FROM `orderdetail`;"

			rows2, err2 := db.Query(detailquery)
			if err2 != nil {
				log.Println(err2)
			}

			for rows2.Next() { // loop orderdetail
				var orderdetail model.OrderDetail

				fmt.Println(order.ID, orderdetail.OrderID)
				if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
					log.Fatal(err.Error())
				} else {
					fmt.Println(orderdetail)
					if order.ID == orderdetail.OrderID {
						orderdetails = append(orderdetails, orderdetail)
					}

				}
			}
			fmt.Println()
			fmt.Println("id", order.ID, "orderdetails", orderdetails)
			transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
			orderdetails = nil
			fmt.Println("trans", transaction.DetailOrder)
		}
		// transaction.DataOrder = append(transaction.DataOrder, order)
		transaction.DataOrder = order
		transactions = append(transactions, transaction)
		transaction.DetailOrder = nil
	}

	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Transactions Success"
		Response.Data = transactions
		sendTransactionSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Transaction Query Error"
		fmt.Print(err)
		sendTransactionErrorResponse(c, Response)
	}

}

func GetFinishedTransaction(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order` WHERE status = 2"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	var order model.Order
	var orders []model.Order
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			if order.Status == 2 {
				orders = append(orders, order)

				detailquery := "SELECT * FROM `orderdetail` WHERE order_id = '" + strconv.Itoa(order.ID) + "';"

				rows2, err2 := db.Query(detailquery)
				if err2 != nil {
					log.Println(err2)
				}

				for rows2.Next() { // loop orderdetail
					var orderdetail model.OrderDetail

					if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
						log.Fatal(err.Error())
					} else {
						if orderdetail.OrderID == order.ID {
							orderdetails = append(orderdetails, orderdetail)
						}
					}
				}

				transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
				orderdetails = nil
		}
		transaction.DataOrder = order
		transactions = append(transactions, transaction)
		transaction.DetailOrder = nil
		}
	}
	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Finished Transaction Success"
		Response.Data = transactions
		sendTransactionSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Finished Transaction Query Error"
		fmt.Print(err)
		sendTransactionErrorResponse(c, Response)
	}
}

func GetTransactionByPizza(c *gin.Context) {
	db := connect()
	defer db.Close()

// 	pizza, _ := strconv.Atoi(c.PostForm("pizza_id"))
	pizza, _ := strconv.Atoi(c.Param("pizza_id"))

	query := "SELECT * FROM `order`"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	var order model.Order
	var orders []model.Order
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)

			detailquery := "SELECT * FROM `orderdetail` WHERE pizza_id='" + strconv.Itoa(pizza) + "';"

			rows2, err2 := db.Query(detailquery)
			if err2 != nil {
				log.Println(err2)
			}

			for rows2.Next() { // loop orderdetail
				var orderdetail model.OrderDetail

				// fmt.Println(order.ID, orderdetail.OrderID)
				if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
					log.Fatal(err.Error())
				} else {
					// fmt.Println(orderdetail)
					if orderdetail.OrderID == order.ID {
						orderdetails = append(orderdetails, orderdetail)
					}
				}
				// fmt.Println("orderdetailID: ", orderdetail.OrderID)
				// fmt.Println("orderID: ", order.ID)
			}

			// fmt.Println("OrderID: ", order.ID, " OrderDetails: ", orderdetails)
			transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
			orderdetails = nil
			// fmt.Println("trans", transaction.DetailOrder)

			// fmt.Println(".")
			// fmt.Println(".")
		}

		// fmt.Println(transactions)

		transaction.DataOrder = order
		transactions = append(transactions, transaction)
		transaction.DetailOrder = nil
	}
	for i := 0; i < len(transactions); i++ {
		// fmt.Println("current: ", i)
		if transactions[i].DetailOrder == nil {
			transactions = RemoveIndex(transactions, i)

			// fmt.Println("index: ", i)
			i = i - 1
		}
	}
	fmt.Println(transactions)
	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Transactions Success"
		Response.Data = transactions
		sendTransactionSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Transaction Query Error"
		fmt.Print(err)
		sendTransactionErrorResponse(c, Response)
	}
}

func GetTransactionByDate(c *gin.Context) {
	db := connect()
	defer db.Close()

	date := c.PostForm("date")

	query := "SELECT * FROM `order`"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	var order model.Order
	var orders []model.Order
	var orderdetails []model.OrderDetail

	for rows.Next() { // loop order
		if err := rows.Scan(&order.ID, &order.CustomerEmail, &order.Waktu,
			&order.Alamat, &order.Status, &order.Rating); err != nil {
			log.Fatal(err.Error())
		} else {

			if strings.Contains(order.Waktu, date) {
				orders = append(orders, order)

				detailquery := "SELECT * FROM `orderdetail`;"

				rows2, err2 := db.Query(detailquery)
				if err2 != nil {
					log.Println(err2)
				}

				for rows2.Next() { // loop orderdetail
					var orderdetail model.OrderDetail

					if err := rows2.Scan(&orderdetail.ID, &orderdetail.PizzaID, &orderdetail.OrderID, &orderdetail.Quantity, &orderdetail.TotalHarga); err != nil {
						log.Fatal(err.Error())
					} else {
						if orderdetail.OrderID == order.ID {
							orderdetails = append(orderdetails, orderdetail)
						}
					}
				}

				transaction.DetailOrder = append(transaction.DetailOrder, orderdetails...)
				orderdetails = nil
			}

		}

		transaction.DataOrder = order
		transactions = append(transactions, transaction)
		transaction.DetailOrder = nil
	}
	// for i := 0; i < len(transactions); i++ {
	// 	// fmt.Println("current: ", i)
	// 	if transactions[i].DetailOrder == nil {
	// 		transactions = RemoveIndex(transactions, i)

	// 		// fmt.Println("index: ", i)
	// 		i = i - 1
	// 	}
	// }

	var Response model.TransactionResponse
	if err == nil {
		Response.Message = "Get Transactions Success"
		Response.Data = transactions
		sendTransactionSuccessResponse(c, Response)
	} else {
		Response.Message = "Get Transaction Query Error"
		fmt.Print(err)
		sendTransactionErrorResponse(c, Response)
	}
}

func RemoveIndex(s []model.Transaction, index int) []model.Transaction {
	return append(s[:index], s[index+1:]...)
}
