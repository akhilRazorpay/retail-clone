package routes

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes, *gin.Engine) {

	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/user/signup", controllers.SignUp())
	incomingRoutes.POST("/admin/product", controllers.AddProduct())
	incomingRoutes.GET("/products", controllers.Products())
	incomingRoutes.GET("/search", controllers.SearchProduct())
}
