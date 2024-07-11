package fiber

import "github.com/gofiber/fiber/v2"

func FiberStart() {
	app := fiber.New()

	app.Post("/calulator", Handler)

	app.Listen(":8080")
}
