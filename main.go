package main

import (
	"app/controllers"
	"app/services"
	"app/models"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	services.ReadConfig()
	models.DatabaseConnect()	
	// models.DatabaseMigrate()
	controllers.Start()
}