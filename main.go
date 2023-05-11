package main

import (
	"net/http"

	"arvo.com/arvo/model"
	"github.com/gin-gonic/gin"
)

// get students list
func getStudents(c *gin.Context) {
	Student := model.GetStudents()
	if Student == nil || len(Student) <= 0 {
		c.JSON(http.StatusNotFound, "Student data not found!")
		return
	}
	c.JSON(http.StatusOK, Student)
}

// get student by id
func getStudent(c *gin.Context) {
	id := c.Param("Id")
	Student := model.GetStudent(id)
	if Student == nil {
		c.JSON(http.StatusNotFound, "Student data not found!")
		return
	}
	c.JSON(http.StatusOK, Student)

}

// add student
func addStudent(c *gin.Context) {
	var newStudent model.Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, "Some thing want wrong!")
		return
	}
	if false == model.AddStudent(newStudent) {
		c.JSON(http.StatusBadRequest, "Some thing want wrong!")
		return
	}
	c.JSON(http.StatusOK, "Student registered!")
}

// update student data
func updateStudent(c *gin.Context) {

	id := c.Param("Id")
	var oldStudent model.Student

	if err := c.BindJSON(&oldStudent); err != nil {
		c.JSON(http.StatusBadRequest, "Some thing want wrong!")
		return
	}

	if false == model.UpdateStudent(oldStudent, id) {
		c.JSON(http.StatusNotFound, "Statudent not found!")
		return
	}
	c.JSON(http.StatusOK, "Student data updated!")

}

//delete student
func deleteStudent(c *gin.Context) {
	id := c.Param("Id")
	if false == model.DeleteStudent(id) {
		c.JSON(http.StatusNotFound, "Student not found!")
		return
	}
	c.JSON(http.StatusOK, "Student data deleted!")
}

// main function
func main() {
	r := gin.Default()
	r.GET("/getStudents", getStudents)
	r.PATCH("/getStudent/:Id", getStudent)
	r.POST("/addStudent", addStudent)
	r.PUT("/Student/:Id", updateStudent)
	r.DELETE("/Student/:Id", deleteStudent)
	r.Run("localhost:808")
}
