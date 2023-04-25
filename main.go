package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/user", getUser)
	r.POST("/user", postUser)
	r.Run(":8080")
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getUser(c *gin.Context) {
	id := c.Query("id")
	user := struct {
		Id string `json:"id"`
		Name string `json:"name"`
	}{id, "Ali"}
	c.JSON(200,gin.H{
		"user": user,
	})
}

func postUser(c *gin.Context) {
	user := struct {
		Id string `json:"id"`
		Name string `json:"name"`
	}{}
	c.BindJSON(&user)

	//сходил в базу данных
	user.Name = "Human"

	c.JSON(200, gin.H{
		"user": user,
	})
}
