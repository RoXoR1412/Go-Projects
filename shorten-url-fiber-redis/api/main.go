package main

import (
	"os"
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/RoXoR1412/shorten-url-fiber-redis/routes"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1",routes.ShortenURL)
}

func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}