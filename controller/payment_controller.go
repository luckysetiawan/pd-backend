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

		fmt.Println(month)
		fmt.Println(payments[i].TotalPembayaran)

		if period == month {
			totalPendapatan += payments[i].TotalPembayaran
			fmt.Println("here")
			fmt.Println(totalPendapatan)
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
