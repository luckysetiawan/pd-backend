package main

import (
	"fmt"
	"time"

	controllers "pd-backend/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("REST API Pizza Delivery")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Melihat seluruh menu
	router.GET("/menu", controllers.GetAllMenus)

	// Login untuk Staff
	router.POST("/login", controllers.Login)
	// Logout untuk Staff
	router.GET("/logout", controllers.Logout)

	// User
	user := router.Group("/user")
	{
		// Update profile user
		user.PUT("/:user_id", controllers.UpdateUser)
	}

	// Admin
	admin := router.Group("/admin")
	{
		// Melihat seluruh member
		admin.GET("/users", controllers.GetAllUsers)
		// Mengubah data user
		admin.PUT("/:user_id", controllers.UpdateUser)
		// Menghapus member
		admin.DELETE("/:user_id", controllers.DeleteUser)
		// Menambahkan menu baru
		admin.POST("/menu", controllers.InsertMenu)
		// Mengubah data menu berdasarkan ID
		admin.PUT("/menu/:menu_id", controllers.UpdateMenu)
		// Menghapus menu
		admin.DELETE("/menu/:menu_id", controllers.DeleteMenu)

		// Melihat seluruh payment
		// admin.GET("/payments", controllers.GetAllPayments)

		// Melihat pendapatan suatu periode
		admin.GET("/payments-period", controllers.GetPaymentForPeriod)
		// Melihat total pizza terjual suatu periode
		admin.GET("/pizza-sold-period", controllers.GetPizzaQuantityForPeriod)
	}

	// Order
	order := router.Group("/order")
	{
		// Melihat seluruh order
		order.GET("/allorder", controllers.GetAllOrders)
		// Melihat Order
		order.GET("/:order_id", controllers.GetOrder)
		// Melihat Status
		order.GET("/status/:order_id", controllers.GetStatus)
		// Membuat order
		order.POST("/create", controllers.InsertOrder)
		// Update order
		order.PUT("/:order_id", controllers.UpdateOrder)
		// Delete order
		order.DELETE("/:order_id", controllers.DeleteOrder)

		// Melihat order yang aktif/dimasak (chef)
		order.GET("/active", controllers.GetActiveOrders)
	}

	transaction := router.Group("/transaction")
	{
		// Menampilkan transaksi yang telah selesai
		transaction.GET("/all", controllers.GetAllTransaction)
		// Menampilkan transaksi yang telah selesai berdasarkan pizza
		transaction.GET("/pizza", controllers.GetTransactionByPizza)
		// Menampilkan transaksi yang telah selesai berdasarkan tanggal
		transaction.GET("/date", controllers.GetTransactionByDate)
	}

	// Payment
	payment := router.Group("/payment")
	{
		// Update payment status menjadi 1 (0=belum dibayar, 1=dibayar)
		payment.PUT("/:order_id", controllers.UpdatePaymentStatus)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
