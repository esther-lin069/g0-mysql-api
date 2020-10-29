package services

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ReturnMessage struct {
	Status     bool       `json:"status"`
	Error_code error      `json:"error"`
	DB_runtime string     `json:"runtime"`
	Result     [][]string `json:"result"`
	RowCount   int        `json:"row_count"`
}

func SqlQuery(f string, c string) ReturnMessage {
	db := CreateDbConn()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		db.Close()
	}()
	st := time.Now()
	rm := ReturnMessage{
		Status:     true,
		Error_code: fmt.Errorf("0"),
		DB_runtime: "",
	}

	//fields := strings.Split(f, ",")
	sql := MakeCode(f, c)
	rows, err := db.Query(sql)
	checkErr(err)

	cols, err := rows.Columns()
	checkErr(err)

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	total_rows := make([][]string, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			rm.Error_code = fmt.Errorf(err.Error())
			rm.DB_runtime = "0"
			rm.Status = false
			return rm
		}

		result := make([]string, len(cols))
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		total_rows = append(total_rows, result)
	}

	rm.DB_runtime = time.Since(st).String()
	rm.Result = total_rows
	rm.RowCount = len(total_rows)

	return rm
}

func MakeCode(fields string, condition string) string {
	var stm string
	if condition == "" {
		stm = "SELECT " + fields + " FROM messages"
	} else {
		stm = "SELECT " + fields + " FROM messages WHERE " + condition
	}

	return stm
}
