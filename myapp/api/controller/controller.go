package controller

import (
	"fmt"
	"main/api/services"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiRespones struct {
	ResultCode   int
	receive_time string
	parameters   string
	apiRunTime   string
}

func GetTime() string {
	tn := time.Now()
	local, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		fmt.Println(err)
	}
	t := tn.In(local)
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return formatted
}

/*資料插入資料庫*/
// func InsertData(c *gin.Context) {
// 	st := time.Now()
// 	r := ResponseCode{
// 		receive_time: time.Now()
// 		db_runtime: services.Insertmessages(),
// 	}
// 	ResponesWithJson(c, ApiRespones{c.Writer.Status(), r, time.Now().Sub(st)})
// }

func FetchAll(c *gin.Context) {
	st := time.Now()
	fields := c.Query("fields")
	result := services.SqlQuery(fields, "")
	api_runTime := time.Since(st)

	response := ApiRespones{
		ResultCode:   c.Writer.Status(),
		receive_time: GetTime(),
		parameters:   fields,
		apiRunTime:   api_runTime.String(),
	}

	ResponesWithJson(c, response, result)
}

func FetchWhere(c *gin.Context) {
	st := time.Now()
	fields := "*"
	condition := c.Request.FormValue("condition")
	result := services.SqlQuery(fields, condition)
	api_runTime := time.Since(st)

	response := ApiRespones{
		ResultCode:   c.Writer.Status(),
		receive_time: GetTime(),
		parameters:   fields,
		apiRunTime:   api_runTime.String(),
	}

	ResponesWithJson(c, response, result)
}

func ResponesWithJson(c *gin.Context, res ApiRespones, rm services.ReturnMessage) {
	c.JSON(res.ResultCode, gin.H{
		"Code":       res.ResultCode,
		"Parameters": res.parameters,
		"Status":     rm.Status,
		"Receive_at": res.receive_time,
		"Error":      rm.Error_code,
		"ApiRunTime": res.apiRunTime,
		"DBRunTime":  rm.DB_runtime,
		"Result":     rm.Result,
	})
}
