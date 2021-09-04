package main

import (
	"log"

	"swagger/database"
	_ "swagger/docs"
	"swagger/models"
	"swagger/routes"
	"github.com/joho/godotenv"
)

// @title Book App
// @version 1.0
// @description This is an API for Book Application

// @BasePath /api
func main() {
	err2 := godotenv.Load()
	if err2 != nil {
		log.Panic(err2)
	}

	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Book{})

	app := routes.New()
	log.Fatal(app.Listen(":3005"))
}
