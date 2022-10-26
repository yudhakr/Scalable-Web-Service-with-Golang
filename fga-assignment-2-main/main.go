package main

import (
	"assignment2/controllers"
	"assignment2/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"

	database.StartDB()

	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)

	fmt.Println("Server is running on port", port)

	r.Run(port)
}
