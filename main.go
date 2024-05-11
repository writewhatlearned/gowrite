package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id       string `json:"id"`
	Male     bool   `json:"male"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InsertData(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": user.Name + " has been inserted successfully!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := "Hello, " + name + "."
		c.String(http.StatusOK, message)
	})

	r.POST("/adduser", InsertData)

	r.Run(":9000")
}
