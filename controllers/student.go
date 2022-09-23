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

func FindStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found"})
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func FindStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found"})
	} else {
		c.JSON(http.StatusOK, student)
	}
}
func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Student not found"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, student)

}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Student removed with success"})
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
