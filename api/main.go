package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/tsivinsky/goenv"
)

func main() {
	err := goenv.Load(Env)
	if err != nil {
		log.Fatal(err)
	}

	err = ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(recover.New())

	Router(app)

	log.Fatal(app.Listen(":5000"))
}
