package middlewares

import (
	"os"

	"backend/api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Authorized() echo.MiddlewareFunc {

	config := middleware.JWTConfig{
		Claims:     &utils.JwtCustomClaimsUser{},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}
	return middleware.JWTWithConfig(config)
}
