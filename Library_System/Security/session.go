package security

import (
	controller "library/Controller"
	errors "library/Errors"
	data "library/JSON"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgtype"
)

type SessionConfig struct {
	Expiration   time.Duration
	CookieSecure bool
}

func SessionCreate(config SessionConfig) {
	app := fiber.New()
	app.Use(SessionMiddleware(config))
	app.Use(cors.New())
	// User islemleri

	app.Post("/login", func(c *fiber.Ctx) error {
		var Data data.ControlUserStr
		err := c.BodyParser(&Data)
		if err != nil {
			errors.Logger(err, 0)
		}
		SetSession(c, Data.Email, Data.Password)
		userId, _ := GetSession(c)

		if userId == 0 {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		} else {
			return c.SendString("Logged in successfully")
		}

	})

	app.Post("/addUser", func(c *fiber.Ctx) error {

		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.AddUserStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.AddUser(Data.Email, Data.Password, Data.Name, Data.Surname, Data.Roles)
			return c.SendString("Used added.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})

	app.Post("/deleteUser", func(c *fiber.Ctx) error {

		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.DeleteUserStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.DeleteUser(Data.ID)
			return c.SendString("Used Deleted.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})
	app.Post("/updateUser", func(c *fiber.Ctx) error {

		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.FullUserStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.UpdateUser(Data.ID, Data.Email, Data.Password, Data.Name, Data.Surname, Data.Roles)
			return c.SendString("Used Updated.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})

	app.Get("/listUser", func(c *fiber.Ctx) error {
		_, Role := GetSession(c)
		if Role == "admin" {

			users := controller.ListUser()

			return c.JSON(users)
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
	})

	//Book
	app.Post("/addBook", func(c *fiber.Ctx) error {
		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.AddBookStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.AddBook(Data.Bookname, Data.Writer, Data.UserID, Data.Deadline)
			return c.SendString("Used added.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})

	app.Post("/deleteBook", func(c *fiber.Ctx) error {
		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.DeleteBookStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.DeleteBook(Data.ID)
			return c.SendString("Used added.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})
	app.Post("/updateBook", func(c *fiber.Ctx) error {
		_, Role := GetSession(c)
		if Role == "admin" {

			var Data data.FullBookStr
			err := c.BodyParser(&Data)
			if err != nil {
				errors.Logger(err, 0)
			}

			controller.UpdateBook(Data.ID, Data.Bookname, Data.Writer, Data.UserID, Data.Deadline)
			return c.SendString("Used added.")
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

	})
	app.Get("/listBook", func(c *fiber.Ctx) error {

		users := controller.ListBooks()

		return c.JSON(users)

	})

	app.Get("/listDead", func(c *fiber.Ctx) error {
		dateStr := c.Query("date")
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			errors.Logger(err, 0)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid date")
		}
		pdate := pgtype.Date{
			Time:             date,
			InfinityModifier: 0,
			Valid:            true,
		}
		users := controller.DeadlineBooks(pdate)

		return c.JSON(users)

	})
	app.Listen(":3000")
}
