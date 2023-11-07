package main

import (
	"log"
	"sesi-7/app/employee"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Meilisearch Services",
	})
	router.Use(cors.New())
	router.Use(logger.New())

	employee.RegisterEmployeeService(router)

	if err := router.Listen(":8080"); err != nil {
		log.Println(err.Error())
	}
}
