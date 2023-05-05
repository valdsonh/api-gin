package main

import (
	"valdson/api-gin/database"
	"valdson/api-gin/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
