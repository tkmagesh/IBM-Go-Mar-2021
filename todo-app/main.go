package main

import (
	"log"

	config "todo-app/config"
	models "todo-app/models"
	routes "todo-app/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildConfig()))

	if err != nil {
		log.Fatalf("Error connecting to database : %s", err.Error())
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})
	r := routes.SetupRouter()

	r.Run()
}
