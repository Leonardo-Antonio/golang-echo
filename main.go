package main

import (
	"log"

	"github.com/Leonardo-Antonio/golang-echo/router"
	"github.com/Leonardo-Antonio/golang-echo/storage/connections"
	"github.com/Leonardo-Antonio/golang-echo/storage/course"
	"github.com/labstack/echo"
)

func main() {
	db := connections.Mysql()
	courseStore := course.New(db)

	e := echo.New()
	router.Course(courseStore, e)

	log.Println("Servidor corriendo en el puerto :8080")

	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Error en el servidor -> %+v\n", err)
	}
}
