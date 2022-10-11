package main

import (
	"Assignment2/database"
	"Assignment2/routers"
)

func main() {
	database.StartDB()
	//var PORT = ":8080"
	var PORT = ":"
	routers.StartServer().Run(PORT)
}
