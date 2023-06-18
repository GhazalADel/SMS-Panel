package server

import (
	database "SMS-panel/database"
	"SMS-panel/handlers"
	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Account
	accountHandler := handlers.NewAccountHandler(db)
	accountRoutes(e, accountHandler)

	// Phonebook
	phonebookHandler := handlers.NewPhonebookHandler(db)
	phonebookRoutes(e, phonebookHandler)

	log.Fatal(e.Start(":8080"))
}
