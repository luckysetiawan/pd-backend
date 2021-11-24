/*
 * Dependencies :
 * gin   : go get github.com/gin-gonic/gin
 * mysql : go get github.com/go-sql-driver/mysql
 * jwt   : go get github.com/dgrijalva/jwt-go
 */
package main

import (
	"fmt"

	controllers "pd-backend/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("REST API Pizza Delivery")

	router := gin.Default()

	// User Registration
	router.POST("/registrasi", controllers.Registrasi)

	// User
	user := router.Group("/user")
	{
		// Update profile user
		user.PUT("/update/:user_id", controllers.UpdateUser)
		// Melihat seluruh menu
		user.GET("/menu", controllers.GetAllPizzas)
	}

	// Admin
	admin := router.Group("/admin")
	{
		// Melihat seluruh member
		admin.GET("/alluser", controllers.GetAllUsers)
		// Mengubah data user
		admin.PUT("/updateuser/:user_id", controllers.UpdateUser)
		// Menghapus member
		admin.DELETE("/deleteuser/:user_id", controllers.DeleteUser)
		// Menambahkan menu baru
		admin.POST("/menu", controllers.InsertPizza)
		// Mengubah data menu berdasarkan ID
		admin.PUT("/menu/:menu_id", controllers.UpdatePizza)
		// Menghapus menu
		admin.DELETE("/deletemenu/:menu_id", controllers.DeletePizza)
	}

	//Order
	order := router.Group("/order")
	{
		// Melihat seluruh order
		order.GET("/allorder", controllers.GetAllOrders)
		// Membuat order
		order.POST("/create", controllers.InsertOrder)
		// Update order
		order.PUT("/update/:order_id", controllers.UpdateOrder)
		//Delete order
		order.DELETE("/delete/:order_id", controllers.DeleteOrder)
	}
	fmt.Println("Connected to port 8080")
	router.Run()
}
