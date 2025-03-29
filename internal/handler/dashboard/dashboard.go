package dashboard

import (
	"github.com/HoneySinghDev/go-echo-rest-api-template/pkg/server"
	"github.com/labstack/echo/v4"
)

func HandleDashboard(_ *server.Server) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, World!"})
	}
}
