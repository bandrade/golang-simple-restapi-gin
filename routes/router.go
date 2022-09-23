package routes

import (
	controller "github.com/bandrade/golang-simple-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controller.ListStudents)
	r.GET("/students/cpf/:cpf", controller.FindStudentByCpf)
	r.GET("/students/:id", controller.FindStudentById)
	r.PUT("/students/:id", controller.EditStudent)
	r.DELETE("/students/:id", controller.DeleteStudent)
	r.POST("/students", controller.CreateStudent)
	r.Run(":5000")
}
