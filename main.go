package main

import (
	"orders/config"
	"orders/controller"
	"os"

	_ "orders/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Api Document Orders
// @version 1.0
// @Description Api Document Order
// @termsOfService http://localhost
// @contact.name pegadaian
// @contact.email pegadaian.id
// @license.name pegadaian 1.0
// @license.url
// @host localhost:9000
// @BasePath /
func main() {
	config.Connect()
	e := echo.New()
	emm := e.Group("/orders")
	emm.POST("", controller.CreateOrder)
	emm.PUT("/:orderId", controller.UpdateOrder)
	emm.GET("", controller.GetAllOrder)
	emm.DELETE("/:orderId", controller.DeleteOrder)
	//route for swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	var port = os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
