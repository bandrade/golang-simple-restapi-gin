package controller

import (
	"net/http"

	"github.com/bandrade/golang-simple-restapi-gin/database"
	"github.com/bandrade/golang-simple-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func ListStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
