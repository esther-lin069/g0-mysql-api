package main

import (
	"main/api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}
