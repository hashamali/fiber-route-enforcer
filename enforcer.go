package enforcer

import (
	"net/http"

	"github.com/gofiber/fiber"
)

// New will create a new route enforcer middleware.
func New() fiber.Handler {
	return func(c *fiber.Ctx) {
		defer func() {
			rawRouteFound := c.Locals(routeFoundKey)
			routeFound, ok := rawRouteFound.(bool)
			if (!ok || !routeFound) && (c.Fasthttp.Response.StatusCode() == http.StatusOK) {
				c.Status(http.StatusNotFound)
			}
		}()

		c.Next()
	}
}

// RouteHandler wraps all routes to mark them as found in a specific context.
func RouteHandler(handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) {
		c.Locals(routeFoundKey, true)
		handler(c)
	}
}

type contextKeyType = string

const routeFoundKey contextKeyType = "routeFound"
