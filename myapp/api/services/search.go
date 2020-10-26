package services

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ReturnMessage struct {
	Status     bool   `json:"status"`
	Error_code error  `json:"error"`
	DB_runtime string `json:"runtime"`
	Result     string `json:"result"`
}

func All(f string) ReturnMessage {
	db := CreateDbConn()
	st := time.Now()

	result := ReturnMessage{
		Status:     true,
		Error_code: fmt.Errorf("0"),
		DB_runtime: time.Since(st).String(),
		Result:     "test",
	}

	return result
}
