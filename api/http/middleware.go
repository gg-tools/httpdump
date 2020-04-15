package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func newAccessLog(logFile *os.File) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logFile})
}
