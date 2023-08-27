package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentDetails struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	CGPA float32 `json:"cgpa"`
}

var class_sheet = []studentDetails{
	{ID: "1", Name: "SUPREETHA", CGPA: 9.88},
	{ID: "2", Name: "DEKUU", CGPA: 8.88},
	{ID: "3", Name: "SPIDERMAN", CGPA: 7.78},
}

func GetStudentDetails(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, class_sheet)
}

func AddStudent(context *gin.Context) {
	var addStudent studentDetails
	if err := context.BindJSON(&addStudent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	class_sheet = append(class_sheet, addStudent)
	context.IndentedJSON(http.StatusOK, addStudent)
}

func GetStudentDetail(context *gin.Context) {
	id := context.Param("id")
	studentDetail, err := GetStudentDetailsByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Student Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, studentDetail)
}

func GetStudentDetailsByID(id string) (*studentDetails, error) {
	for i, s := range class_sheet {
		if s.ID == id {
			return &class_sheet[i], nil
		}
	}
	return nil, errors.New("student not found")
}

func main() {
	routes := gin.Default()
	routes.GET("/sheets", GetStudentDetails)
	routes.POST("/sheet", AddStudent)
	routes.GET("/sheets/:id", GetStudentDetail)
	routes.Run("localhost:8080")
}
