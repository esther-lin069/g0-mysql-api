package routes

import (
	"github.com/gin-gonic/gin"

	"main/api/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/insert", controller.InsertData)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "yes"})
	})

	return r
}
