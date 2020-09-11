package main

import (
	"log"

	"github.com/labstack/echo/middleware"

	"github.com/Leonardo-Antonio/golang-echo/certificates/authorization"

	"github.com/labstack/echo"

	"github.com/Leonardo-Antonio/golang-echo/router"
	"github.com/Leonardo-Antonio/golang-echo/storage/connections"
	"github.com/Leonardo-Antonio/golang-echo/storage/course"
	"github.com/Leonardo-Antonio/golang-echo/storage/user"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados -> %v", err)
	}

	db := connections.Mysql()
	courseStore := course.New(db)
	UserStore := user.New(db)

	e := echo.New()
	e.Use(middleware.CORS())

	router.Course(courseStore, e)
	router.User(UserStore, e)

	log.Println("Servidor corriendo en el puerto :8080")

	err = e.Start(":8080")
	if err != nil {
		log.Fatalf("Error en el servidor -> %+v\n", err)
	}
}
