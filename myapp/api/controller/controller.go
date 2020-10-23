package controller

import (
	"main/api/services"

	"github.com/gin-gonic/gin"
)

type ResponseCode struct {
	receive_time string
	parameters   interface{}
}

/*資料插入資料庫*/
func InsertData(c *gin.Context) {
	t := services.Insertmysql()
	c.JSON(200, gin.H{
		"test": t})
}
