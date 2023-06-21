package router

import (
	res "payment/utils/response"

	"payment/handler"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		res.Response(c, 200, "", "Hello There!")
	})

	app.GET("/test", handler.Test)
	// app.POST("/add", handler.AddOneEmp)
	app.GET("/getall", handler.GetEmployees)
	app.GET("/getone/:name", handler.GetOneEmployee)
	app.PUT("/update", handler.UpdateEmployee)
	app.DELETE("/delete/:name", handler.DeleteOneEmployee)

	ProjectModules(app)
}
