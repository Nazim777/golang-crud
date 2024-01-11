package main

import (
	"user-api/database"
	"user-api/routes"
)

func main() {

	database.ConnectDatabase()
	routes.InitilizeRouter()
}

