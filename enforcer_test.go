package enforcer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber"
)

func TestEnforcerNoRoutes(t *testing.T) {
	a := fiber.New()
	a.Use(New())

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, err := a.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		t.Fatal("should receive 404")
	}
}

func TestEnforcerRoutes(t *testing.T) {
	route := "/test"

	a := fiber.New()
	a.Use(New())
	a.Get(route, RouteHandler(func(c *fiber.Ctx) {
		c.Status(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, route, nil)
	resp, err := a.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		t.Fatal("should receive 204")
	}
}
