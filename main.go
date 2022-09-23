package main

import (
	"github.com/bandrade/golang-simple-restapi-gin/database"
	"github.com/bandrade/golang-simple-restapi-gin/routes"
)

func main() {
	database.ConnectWithDB()
	routes.HandleRequests()
}
