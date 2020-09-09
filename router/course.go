package router

import (
	"github.com/Leonardo-Antonio/golang-echo/handler"
	"github.com/Leonardo-Antonio/golang-echo/middleware"
	"github.com/labstack/echo"

	"github.com/Leonardo-Antonio/golang-echo/storage"
)

// Course -> endpoints
func Course(s storage.Course, c *echo.Echo) {
	courseHandler := handler.NewCourse(s)

	courses := c.Group("/v1/courses")
	courses.Use(middleware.Authorization)
	courses.POST("", courseHandler.Create)
	courses.GET("", courseHandler.GetAll)
	courses.GET("/:id", courseHandler.GetByID)
	courses.PUT("/:id", courseHandler.Update)
	courses.DELETE("/:id", courseHandler.Delete)
}
