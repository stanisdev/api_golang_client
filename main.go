package main

import (
	"app/controllers"
	"app/services"
)

func main() {
	controllers.Start()
	services.ReadConfig()
	services.DatabaseConnect()
}