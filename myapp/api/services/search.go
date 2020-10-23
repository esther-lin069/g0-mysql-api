package services

import (
	_ "github.com/go-sql-driver/mysql"
)

func Query() {
	db := CreateDbConn()
}
