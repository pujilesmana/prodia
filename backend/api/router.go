package api

import (
	"backend/api/middlewares"
	"backend/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all V1 API with routing path

type Controller struct {
	User *user.UserController
}

func RegisterPath(e *echo.Echo, c *Controller) {

	if c.User == nil {
		panic("user controller cannot be nil")
	}

	userV1 := e.Group("api/v1/users")
	userV1.GET("", c.User.GetCurrentUser, middlewares.Authorized())
	userV1.POST("", c.User.CreateNewUser)
	userV1.POST("/login", c.User.Login)
	userV1.PUT("/", c.User.UpdateUser, middlewares.Authorized())
}
