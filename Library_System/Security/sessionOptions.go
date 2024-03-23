package security

import (
	user "library/Controller"
	errors "library/Errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SessionMiddleware(config SessionConfig) func(*fiber.Ctx) error {
	store := session.New(session.Config{
		Expiration:   config.Expiration,
		CookieSecure: config.CookieSecure,
	})
	return func(c *fiber.Ctx) error {

		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		err = sess.Save()
		if err != nil {
			return err
		}

		c.Locals("sessionStore", store)
		return c.Next()
	}
}
func GetSession(c *fiber.Ctx) (int32, string) {

	sessionStore := c.Locals("sessionStore").(*session.Store)

	session, _ := sessionStore.Get(c)
	id := session.Get("id")
	role := session.Get("role")
	if id != nil {
		idInt := id.(int32)
		roleStr := role.(string)
		return idInt, roleStr
	} else {
		return 0, ""
	}

}

func SetSession(c *fiber.Ctx, email, pass string) {
	sessionStore := c.Locals("sessionStore").(*session.Store)

	session, _ := sessionStore.Get(c)

	sessionCred := user.ControlUser(email, pass)

	session.Set("id", sessionCred.ID)
	session.Set("role", sessionCred.Roles.String)
	err := session.Save()
	if err != nil {
		errors.Logger(err, 0)
	}
}
