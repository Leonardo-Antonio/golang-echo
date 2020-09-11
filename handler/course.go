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
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "POST"}
		response := newResponse(ERROR, "No tiene la estructura correcta", nil, i)
		return cx.JSON(i.Code, response)
	}

	err = c.storage.Create(data)
	if errors.Is(err, storage.ErrorRowNull) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "POST"}
		response := newResponse(ERROR, "No se aceptan datos nulos", nil, i)
		return cx.JSON(i.Code, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "POST"}
		response := newResponse(ERROR, "No se afectaron las filas deseadas", nil, i)
		return cx.JSON(i.Code, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: cx.Path(), Method: "POST"}
		response := newResponse(ERROR, "Ha ocurrido un error", nil, i)
		return cx.JSON(i.Code, response)
	}

	i := info{Code: http.StatusCreated, Path: cx.Path(), Method: "POST"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return cx.JSON(i.Code, response)
}

// GetAll -> method of class Course
func (c *Course) GetAll(cx echo.Context) error {
	courses, err := c.storage.GetAll()
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: cx.Path(), Method: "GET"}
		response := newResponse(ERROR, "No se pudo obtener la información", nil, i)
		return response.JSONXML(cx, response)
	}
	i := info{Code: http.StatusOK, Path: cx.Path(), Method: "GET"}
	response := newResponse(MESSAGE, "OK", courses, i)
	return response.JSONXML(cx, response)
}

// GetByID -> method of class Course
func (c *Course) GetByID(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "GET"}
		response := newResponse(ERROR, "El id debe ser un número entero", nil, i)
		return response.JSONXML(cx, response)
	}

	course, err := c.storage.GetByID(ID)
	if errors.Is(err, storage.ErrorNotExistCourse) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "GET"}
		response := newResponse(ERROR, "El curso no existe", nil, i)
		return response.JSONXML(cx, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: cx.Path(), Method: "GET"}
		response := newResponse(ERROR, "Ha ocurrido un error", nil, i)
		return response.JSONXML(cx, response)
	}

	i := info{Code: http.StatusOK, Path: cx.Path(), Method: "GET"}
	response := newResponse(MESSAGE, "OK", course, i)
	return response.JSONXML(cx, response)
}

// Update -> method of class Course
func (c *Course) Update(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "PUT"}
		response := newResponse(ERROR, "El id debe ser un número entero", nil, i)
		return response.JSONXML(cx, response)
	}

	data := model.Course{}
	err = cx.Bind(&data)
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "PUT"}
		response := newResponse(ERROR, "La estructura no es correcta", nil, i)
		return response.JSONXML(cx, response)
	}

	err = c.storage.Update(ID, data)
	if errors.Is(err, storage.ErrorRowNull) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "PUT"}
		response := newResponse(ERROR, "No se aceptan datos nulos", nil, i)
		return response.JSONXML(cx, response)
	}
	if errors.Is(err, storage.ErrorNotExistCourse) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "PUT"}
		response := newResponse(ERROR, "No existe el id: "+cx.Param("id")+" o ya esta actualizado", nil, i)
		return response.JSONXML(cx, response) // error de idepotencia
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: cx.Path(), Method: "PUT"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(cx, response)
	}

	i := info{Code: http.StatusOK, Path: cx.Path(), Method: "PUT"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return response.JSONXML(cx, response)
}

// Delete -> method of class Course
func (c *Course) Delete(cx echo.Context) error {
	ID, err := strconv.Atoi(cx.Param("id"))
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "El id debe ser un número entero", nil, i)
		return response.JSONXML(cx, response)
	}

	err = c.storage.Delete(ID)
	if errors.Is(err, storage.ErrorNotExistCourse) {
		i := info{Code: http.StatusBadRequest, Path: cx.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "No existe el id: "+cx.Param("id"), nil, i)
		return response.JSONXML(cx, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: cx.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(cx, response)
	}

	i := info{Code: http.StatusOK, Path: cx.Path(), Method: "DELETE"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return response.JSONXML(cx, response)
}
