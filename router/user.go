package router

import (
	"github.com/Leonardo-Antonio/golang-echo/handler"
	"github.com/Leonardo-Antonio/golang-echo/storage"
	"github.com/labstack/echo"
)

// User -> endpoints of users
func User(s storage.User, e *echo.Echo) {
	user := handler.NewUser(s)
	users := e.Group("/v1/users")

	users.DELETE("/:id", user.Delete)
	users.POST("", user.LogIn)
	users.POST("/new", user.SignIn)
	users.PUT("/:id", user.Update)

}
