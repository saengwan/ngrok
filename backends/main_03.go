package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employee", EmployeeController.GetEmployee)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID
	router.POST("/employee", EmployeeController.PostEmployee)       //POST
	router.PUT("/employee", EmployeeController.PutEmployee)         //PUT
	router.DELETE("/employee", EmployeeController.DeleteEmployee)   //DELETE
	router.Run()                                                    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
