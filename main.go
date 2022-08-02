package main

import (
	"retail-clone/controllers"
	"retail-clone/models"

	"github.com/gin-gonic/gin"
)

func main() {
	A := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	A.GET("/products", controllers.Findproducts)
	A.GET("/products/:id", controllers.Findproduct)
	A.POST("/products", controllers.Createproduct)
	A.PATCH("/products/:id", controllers.Updateproduct)
	A.DELETE("/products/:id", controllers.Deleteproduct)

	// Run the server
	A.Run()
}
