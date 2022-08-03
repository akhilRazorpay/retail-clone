package main

import (
	"fmt"
	"log"
	"retail-clone/controllers"
	"retail-clone/database"
	"retail-clone/middleware"
	"retail-clone/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8000"
	// }
	app := controllers.NewApplication(database.ProductData(database.Client, "products"), database.UserData(database.Client, "users"))

	A := gin.New()
	A.Use(gin.Logger())
	routes.UserRoutes(A)
	A.Use(middleware.Authentication())
	A.GET("/addtocart", app.AddToCart())
	A.GET("removeitem", app.RemoveItem())
	A.GET("/cart checkout", app.BuyFromCart())
	A.GET("/instantbuy", app.InstantBuy())

	// Run the server
	log.Fatal(A.Run())

	fmt.Println("ok")
}
