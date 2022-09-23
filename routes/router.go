package routes

import (
	controller "github.com/bandrade/golang-simple-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controller.ListStudents)
	r.POST("/students", controller.CreateStudent)
	r.Run(":5000")
}
