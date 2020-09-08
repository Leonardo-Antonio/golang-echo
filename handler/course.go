package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/golang-echo/model"
	"github.com/Leonardo-Antonio/golang-echo/storage"
	"github.com/labstack/echo"
)

// Course -> class (handler)
type Course struct {
	storage storage.Course
}

// NewCourse -> method constructor
func NewCourse(s storage.Course) *Course {
	return &Course{s}
}

// Create -> method of class Course
func (c *Course) Create(cx echo.Context) error {
	data := model.Course{}
	err := cx.Bind(&data)
	if err != nil {
		response := newResponse(ERROR, "No tiene la estructura correcta", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}

	err = c.storage.Create(data)
	if errors.Is(err, storage.ErrorRowNull) {
		response := newResponse(ERROR, "No se aceptan datos nulos", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		response := newResponse(ERROR, "No se afectaron las filas deseadas", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error", nil)
		return cx.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", nil)
	return cx.JSON(http.StatusCreated, response)
}

// GetAll -> method of class Course
func (c *Course) GetAll(cx echo.Context) error {
	courses, err := c.storage.GetAll()
	if err != nil {
		response := newResponse(ERROR, "No se pudo obtener la información", nil)
		return cx.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(MESSAGE, "OK", courses)
	return cx.JSON(http.StatusOK, response)
}

// GetByID -> method of class Course
func (c *Course) GetByID(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		response := newResponse(ERROR, "El id debe ser un número entero", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}

	course, err := c.storage.GetByID(ID)
	if errors.Is(err, storage.ErrorNotExistCourse) {
		response := newResponse(ERROR, "El curso no existe", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error", nil)
		return cx.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", course)
	return cx.JSON(http.StatusOK, response)

}

// Update -> method of class Course
func (c *Course) Update(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		response := newResponse(ERROR, "El id debe ser un número entero", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}

	data := model.Course{}
	err = cx.Bind(&data)
	if err != nil {
		response := newResponse(ERROR, "La estructura no es correcta", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}

	err = c.storage.Update(ID, data)
	if errors.Is(err, storage.ErrorRowNull) {
		response := newResponse(ERROR, "No se aceptan datos nulos", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if errors.Is(err, storage.ErrorNotExistCourse) {
		response := newResponse(ERROR, "No existe el id: "+cx.Param("id"), nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return cx.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", nil)
	return cx.JSON(http.StatusOK, response)

}

// Delete -> method of class Course
func (c *Course) Delete(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		response := newResponse(ERROR, "El id debe ser un número entero", nil)
		return cx.JSON(http.StatusBadRequest, response)
	}

	err = c.storage.Delete(ID)
	if errors.Is(err, storage.ErrorNotExistCourse) {
		response := newResponse(ERROR, "No existe el id: "+cx.Param("id"), nil)
		return cx.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return cx.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", nil)
	return cx.JSON(http.StatusInternalServerError, response)
}
