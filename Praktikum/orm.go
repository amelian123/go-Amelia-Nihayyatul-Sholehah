package main

import (
	"orm/config"
	"orm/routes"
)

func main() {
	//start the server, and log if it fails

	config.InitDB()
	config.InitialMigration()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
