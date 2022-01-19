package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getResult(c *gin.Context) {
	marks := c.Query("marks")
	intMarks, err := strconv.Atoi(marks)
	//fmt.Println("Enter the student marks:")
	//var marks int
	if err != nil {
		fmt.Println(err)
	}
	c.String(200, "##### Student Marks Result Portal #####")
	//fmt.Scanln(&marks)
	if intMarks >= 500 {
		c.String(200, "\nDistinction")
	} else if intMarks >= 400 {
		c.String(200, "\nFirst Class")
	} else if intMarks >= 300 {
		c.String(200, "\nSecond Class")
	} else if intMarks >= 200 {
		c.String(200, "\nThrid Class")
	} else {
		c.String(200, "\nFailed! Sorry")
	}
	c.String(200, "\n")
	c.String(200, "\n##### Result Query Success #####")
}

func main() {
	router := gin.Default()
	router.GET("/get", getResult)
	router.Run("localhost:8080")
}
