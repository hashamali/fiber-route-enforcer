# Fiber Route Enforcer
[![godoc](https://godoc.org/github.com/hashamali/fiber-route-enforcer?status.svg)](http://godoc.org/github.com/hashamali/fiber-route-enforcer)
[![tests](https://img.shields.io/github/workflow/status/hashamali/fiber-route-enforcer/tests?label=tests&style=flat-square)](https://github.com/hashamali/fiber-route-enforcer/actions?query=workflow%3Atests)
[![sec](https://img.shields.io/github/workflow/status/hashamali/fiber-route-enforcer/security?label=security&style=flat-square)](https://github.com/hashamali/fiber-route-enforcer/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/fiber-route-enforcer)](https://goreportcard.com/report/github.com/hashamali/fiber-route-enforcer)
[![license](https://badgen.net/github/license/hashamali/fiber-route-enforcer)](https://opensource.org/licenses/MIT)

[Fiber](https://github.com/gofiber/fiber) has a curious behaviour of of allowing non-matched routes to return 200 *if* they happen to "match" a middleware. This can happen quite often if you use any of the top level middlewares such as [logger](https://github.com/gofiber/logger) or [requestid](https://github.com/gofiber/requestid).

This packed forces an request that hasn't been matched by a route to return 404 instead.

#### Usage

Example app. Here, only `/ping` will return a 200. Anything else will return 404. However, the logger will still log the appropriate status code throughout.

```
import (
    "github.com/gofiber/fiber"
    "github.com/gofiber/logger"

    enforcer "github.com/hashamali/fiber-route-enforcer"
)

func main() {
    app := fiber.New()
    app.Use(logger.New())
    app.Use(enforcer.New())
    app.Get("/ping", enforcer.RouteHandler(
        func(c *fiber.Ctx) {
            c.Send(requestid.Get(c))
        },
    ))

    err := app.Listen(8080)
    if err != nil {
        panic(err)
    }
}
```