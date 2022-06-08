package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New()
	err := app.Listen("3000")
	if err != nil {
		log.Fatalln("Bad server")
	}
}
