package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luccasniccolas/monitor/config"
	"github.com/luccasniccolas/monitor/database"
	"github.com/luccasniccolas/monitor/handlers"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error al cargar las variables de configuracion")
	}

	database.DB, err = database.ConnectDatabase(&cfg)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos")
	}
	defer database.DB.Close()

	app := fiber.New()
	api := app.Group("/api")
	auth := api.Group("/auth") // Ruta para metodos de autenticación

	auth.Post("signup", handlers.SignUp)

	app.Listen(":8080")

}
