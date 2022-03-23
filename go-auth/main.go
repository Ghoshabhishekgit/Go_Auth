package main

import (
	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	_, err := gorm.Open(mysql.Open("root@Sunlight@123@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("could not connect to to the database")
	}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
