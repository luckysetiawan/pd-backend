package main

import (
	"fmt"
	"time"

	controllers "pd-backend/controller"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("REST API Pizza Delivery")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH","DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// User Registration
	router.POST("/registrasi", controllers.Registrasi)

	// User
	user := router.Group("/user")
	{
		// Update profile user
		user.PUT("/update/:user_id", controllers.UpdateUser)
		// Melihat seluruh menu
		user.GET("/menu", controllers.GetAllMenus)
	}

	// Admin
	admin := router.Group("/admin")
	{
		// Login Staff
		admin.POST("/login", controllers.Login)
		// Logout Staff
		admin.GET("/logout", controllers.Logout)

		// Melihat seluruh member
		admin.GET("/alluser", controllers.GetAllUsers)
		// Mengubah data user
		admin.PUT("/updateuser/:user_id", controllers.UpdateUser)
		// Menghapus member
		admin.DELETE("/deleteuser/:user_id", controllers.DeleteUser)
		// Menambahkan menu baru
		admin.POST("/menu", controllers.InsertMenu)
		// Mengubah data menu berdasarkan ID
		admin.PUT("/menu/:menu_id", controllers.UpdateMenu)
		// Menghapus menu
		admin.DELETE("/deletemenu/:menu_id", controllers.DeleteMenu)
	}

	//Order
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
		order.PUT("/update/:order_id", controllers.UpdateOrder)
		// Delete order
		order.DELETE("/delete/:order_id", controllers.DeleteOrder)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
