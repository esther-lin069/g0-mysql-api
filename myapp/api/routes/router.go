package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/api/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	//r.GET("/insert", controller.InsertData)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/all", controller.FetchAll)

	return r
}
