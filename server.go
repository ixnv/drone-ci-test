package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	list := List{}

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			ProjectName string `json:"project"`
		}{
			ProjectName: "single linked list",
		})
	})

	router.GET("/append/:value", func(c *gin.Context) {
		value := c.Param("value")

		node := Node{value, nil}

		list.Append(&node)

		c.JSON(http.StatusOK, list)
	})

	router.GET("/display", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})

	if err := router.Run(":8080"); err != nil {
		panic("port is busy")
	}
}
