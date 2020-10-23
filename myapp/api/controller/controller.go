package controller

import (
	"encoding/json"
	"main/api/services"

	"github.com/gin-gonic/gin"
)

type ResponseCode struct {
	receive_time string
	parameters   interface{}
	status       string
	error_code   string
	api_runtime  string
	db_runtime   string
}

type ApiRespones struct {
	ResultCode    int
	ResultMessage interface{}
}

/*資料插入資料庫*/
func InsertData(c *gin.Context) {
	r := ResponseCode{
		db_runtime: services.Insertmessages(),
	}
	ResponesWithJson(c, ApiRespones{200, r})
}

func ResponesWithJson(c *gin.Context, res ApiRespones) {
	response, _ := json.Marshal(res)
	c.JSON(res.ResultCode, gin.H{
		"code":   res.ResultCode,
		"Result": response,
	})
}
