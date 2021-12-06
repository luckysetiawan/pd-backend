package controllers

import (
	"fmt"
	"log"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

func GetAllPayments(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM payment"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	var payment model.Payment
	var payments []model.Payment
	for rows.Next() {
		if err := rows.Scan(&payment.ID, &payment.OrderID, &payment.StatusPembayaran,
			&payment.TotalPembayaran, &payment.WaktuPembayaran); err != nil {
			log.Fatal(err.Error())
		} else {
			payments = append(payments, payment)
		}

	}
	var response model.PaymentResponse
	if err == nil {
		response.Message = "Get Payment Success"
		response.Data = payments
		sendPaymentSuccessresponse(c, response)
	} else {
		response.Message = "Get Payment Error"
		sendPaymentErrorResponse(c, response)
	}
}

func GetPaymentForPeriod(c *gin.Context) {
	db := connect()
	defer db.Close()

	period := c.PostForm("period") //Period yang digunakan adalah month/bulan
	totalPendapatan := 0

	query := "SELECT * FROM payment"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	var payment model.Payment
	var payments []model.Payment
	for rows.Next() {
		if err := rows.Scan(&payment.ID, &payment.OrderID, &payment.StatusPembayaran,
			&payment.TotalPembayaran, &payment.WaktuPembayaran); err != nil {
			log.Fatal(err.Error())
		} else {
			payments = append(payments, payment)
		}
	}

	totalPendapatan = 0
	for i := 0; i < len(payments); i += 1 {
		month := ""
		month = string(payments[i].WaktuPembayaran[5])
		month += string(payments[i].WaktuPembayaran[6])

		// fmt.Println(month)
		// fmt.Println(payments[i].TotalPembayaran)

		if period == month {
			totalPendapatan += payments[i].TotalPembayaran
			// fmt.Println(totalPendapatan)
		}

	}

	var response model.PendapatanResponse
	if err == nil {
		response.Message = "Get Pembayaran Untuk Periode " + period + " Success"
		response.Total = totalPendapatan
		sendGetPendapatanPeriodSuccessResponse(c, response)
	} else {
		response.Message = "Get Menu Query Error"
		sendGetPendapatanPeriodErrorResponse(c, response)
	}
}

// Update Payment Status
func UpdatePaymentStatus(c *gin.Context) {
	db := connect()
	defer db.Close()

	orderId := c.Param("order_id")

	rows, _ := db.Query("SELECT * FROM payment WHERE order_id='" + orderId + "'")
	var payment model.Payment
	for rows.Next() {
		if err := rows.Scan(&payment.ID, &payment.OrderID, &payment.StatusPembayaran,
			&payment.TotalPembayaran, &payment.WaktuPembayaran); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Mengubah status pembayaran menjadi "dibayar(1)"
	status_pembayaran := 1

	_, errQuery := db.Exec("UPDATE payment SET status_pembayaran = ? WHERE order_id=?",
		status_pembayaran,
		orderId,
	)

	var response model.PaymentResponse
	if errQuery == nil {
		response.Message = "Update Payment Status Success"
		sendPaymentSuccessresponse(c, response)
	} else {
		response.Message = "Update Payment Status Error"
		sendPaymentErrorResponse(c, response)
	}
}

func GetPizzaQuantityForPeriod(c *gin.Context) {
	db := connect()
	defer db.Close()

	// Period yang digunakan adalah month/bulan
	period := c.PostForm("period")
	totalPizzaTerjual := 0

	query_1 := "SELECT * FROM payment"

	rows_1, err := db.Query(query_1)
	if err != nil {
		log.Println(err)
	}
	var payment model.Payment
	var payments []model.Payment
	for rows_1.Next() {
		if err := rows_1.Scan(&payment.ID, &payment.OrderID, &payment.StatusPembayaran,
			&payment.TotalPembayaran, &payment.WaktuPembayaran); err != nil {
			log.Fatal(err.Error())
		} else {
			payments = append(payments, payment)
		}
	}

	query_2 := "SELECT * FROM orderdetail"

	rows_2, err := db.Query(query_2)
	if err != nil {
		log.Println(err)
	}
	var orderDetail model.OrderDetail
	var orderDetails []model.OrderDetail
	for rows_2.Next() {
		if err := rows_2.Scan(&orderDetail.ID, &orderDetail.PizzaID, &orderDetail.OrderID,
			&orderDetail.Quantity, &orderDetail.TotalHarga); err != nil {
			log.Fatal(err.Error())
		} else {
			orderDetails = append(orderDetails, orderDetail)
		}
	}

	totalPizzaTerjual = 0
	for i := 0; i < len(payments); i += 1 {
		month := ""
		month = string(payments[i].WaktuPembayaran[5])
		month += string(payments[i].WaktuPembayaran[6])

		fmt.Println(month)
		fmt.Print("orderID payments: ")
		fmt.Println(payments[i].OrderID)
		fmt.Print("orderID orderDetails: ")
		fmt.Println(orderDetails[i].OrderID)
		fmt.Print("pizza quantity: ")
		fmt.Println(orderDetails[i].Quantity)
		fmt.Print("status pembayaran: ")
		fmt.Println(payments[i].StatusPembayaran)

		// Jika status pembayaran masih 0 (belum dibayar), total pizza terjual tidak akan bertambah
		if period == month && payments[i].OrderID == orderDetails[i].OrderID && payments[i].StatusPembayaran == 1 {
			totalPizzaTerjual += orderDetails[i].Quantity
			fmt.Print("total pizza terjual: ")
			fmt.Println(totalPizzaTerjual)
		}

	}

	var response model.PizzaTerjualResponse
	if err == nil {
		response.Message = "Get Pizza Terjual Untuk Periode " + period + " Success"
		response.Total = totalPizzaTerjual
		sendGetPizzaTerjualPeriodSuccessResponse(c, response)
	} else {
		response.Message = "Get Pizza Terjual Query Error"
		sendGetPizzaTerjualPeriodErrorResponse(c, response)
	}
}
