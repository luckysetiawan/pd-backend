package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Get All Orders
func GetAllOrders(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM `order` o JOIN users u on o.idCustomer=u.id JOIN OrderDetail od ON o.idOrderDetail=od.id JOIN payment p on o.invoice=p.invoice;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order model.Order
	var orders []model.Order
	var user model.User
	var users []model.User
	var orderDetail model.OrderDetail
	var orderDetails []model.OrderDetail
	var payment model.Payment
	var payments []model.Payment
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.IdCustomer, &order.IdOrderDetail, &order.Invoice,
			&order.Waktu, &order.Alamat, &order.Status, &user.ID, &user.Nama, &user.Email, &user.NoTelp, &user.Password,
			&orderDetail.ID, &orderDetail.Pizza, &orderDetail.Rating,
			&orderDetail.Quantity, &orderDetail.TotalPesanan,
			&payment.Invoice, &payment.StatusPembayaran, &payment.TotalHarga); err != nil {
			log.Fatal(err.Error)
		} else {
			orders = append(orders, order)
			users = append(users, user)
			orderDetails = append(orderDetails, orderDetail)
			payments = append(payments, payment)
		}
	}

	var OrderFullResponse model.OrderFullResponse
	if err == nil {
		OrderFullResponse.Message = "Get Order Success"
		OrderFullResponse.DataOrder = orders
		OrderFullResponse.DataUser = users
		OrderFullResponse.DataOrderDetail = orderDetails
		OrderFullResponse.DataPayment = payments
		sendFullOrderSuccessresponse(c, OrderFullResponse)
	} else {
		OrderFullResponse.Message = "Get Order Query Error"
		fmt.Print(err)
		sendFullOrderErrorResponse(c, OrderFullResponse)
	}
}

// Insert Order
func InsertOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID, _ := strconv.Atoi(c.PostForm("id"))
	IdCustomer, _ := strconv.Atoi(c.PostForm("idCustomer"))
	IdOrderDetail, _ := strconv.Atoi(c.PostForm("idOrderDetail"))
	Invoice := "P-" + string(time.Now().Format("2006-01-02 15:04:05"))
	Waktu := time.Now()
	Alamat := c.PostForm("alamat")
	Status, _ := strconv.Atoi(c.PostForm("status"))
	_, errQuery := db.Exec("INSERT INTO `order`(id, idCustomer, idOrderDetail, invoice, waktu, alamat, status) values (?,?,?,?,?,?,?)",
		ID,
		IdCustomer,
		IdOrderDetail,
		Invoice,
		Waktu,
		Alamat,
		Status,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Insert Order Success"
		fmt.Print(Invoice)
		sendOrderSuccessresponse(c, response)
	} else {
		response.Message = "Insert Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorResponse(c, response)
	}
}

// Update Order
func UpdateOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	ID := c.Param("order_id")
	IdCustomer := c.PostForm("idCustomer")
	IdOrderDetail := c.PostForm("idOrderDetail")
	Invoice := c.PostForm("invoice")
	Waktu := c.PostForm("waktu")
	Alamat := c.PostForm("alamat")
	Status := c.PostForm("status")

	rows, _ := db.Query("SELECT * FROM `order` WHERE id='" + ID + "'")
	var order model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.IdCustomer, &order.IdOrderDetail, &order.Invoice, &order.Waktu, &order.Alamat, &order.Status); err != nil {
			log.Fatal(err.Error())
		}
	}

	// Jika kosong dimasukkan nilai lama
	if ID == "" {
		ID = strconv.Itoa(order.ID)
	}

	if IdCustomer == "" {
		IdCustomer = strconv.Itoa(order.IdCustomer)
	}

	if IdOrderDetail == "" {
		IdOrderDetail = strconv.Itoa(order.IdOrderDetail)
	}

	if Invoice == "" {
		Invoice = order.Invoice
	}

	if Waktu == "" {
		Waktu = order.Waktu
	}

	if Alamat == "" {
		Alamat = order.Alamat
	}

	if Status == "" {
		Status = strconv.Itoa(order.Status)
	}
	_, errQuery := db.Exec("UPDATE `order` SET idCustomer = ?, idOrderDetail = ?, invoice = ?, Waktu = ?, Alamat = ?, Status = ? WHERE id=?",
		IdCustomer,
		IdOrderDetail,
		Invoice,
		Waktu,
		Alamat,
		Status,
		ID,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Update Order Success"
		sendOrderSuccessresponse(c, response)
	} else {
		response.Message = "Update Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorResponse(c, response)
	}
}

// Delete Order...
func DeleteOrder(c *gin.Context) {
	db := connect()
	defer db.Close()

	orderId := c.Param("order_id")

	_, errQuery := db.Exec("DELETE FROM `order` WHERE id=?",
		orderId,
	)

	var response model.OrderResponse
	if errQuery == nil {
		response.Message = "Delete Order Success"
		sendOrderSuccessresponse(c, response)
	} else {
		response.Message = "Delete Order Failed Error"
		fmt.Print(errQuery)
		sendOrderErrorResponse(c, response)
	}
}

func sendOrderSuccessresponse(c *gin.Context, ur model.OrderResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendFullOrderSuccessresponse(c *gin.Context, ur model.OrderFullResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendOrderErrorResponse(c *gin.Context, ur model.OrderResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func sendFullOrderErrorResponse(c *gin.Context, ur model.OrderFullResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
