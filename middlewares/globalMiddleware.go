package middlewares

import (
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func GlobalMiddlewares(ec *echo.Echo) {
	ec.Use(m.LoggerWithConfig(m.LoggerConfig{Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n"}))

	ec.Use(m.AddTrailingSlash())
}
