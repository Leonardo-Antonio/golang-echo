package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/golang-echo/certificates/authorization"
	"github.com/Leonardo-Antonio/golang-echo/model"

	"github.com/labstack/echo"

	"github.com/Leonardo-Antonio/golang-echo/storage"
)

// User -> class handler
type User struct {
	storage storage.User
}

// NewUser -> method constructor
func NewUser(s storage.User) *User {
	return &User{s}
}

// SignIn -> method of the class User (handler)
func (u *User) SignIn(c echo.Context) error {
	data := model.User{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(ERROR, "La estructura no es la correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = u.storage.SignIn(data)
	if errors.Is(err, storage.ErrorRowNull) {
		response := newResponse(ERROR, "No se aceptan datos nulos", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		response := newResponse(ERROR, "No se pudo crear el nuevo usuario", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(MESSAGE, "OK", nil)
	return c.JSON(http.StatusOK, response)
}

// LogIn -> method of the class User (handler)
func (u *User) LogIn(c echo.Context) error {
	data := model.User{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(ERROR, "La estructura no es la correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err = u.storage.LogIn(data)
	if errors.Is(err, storage.ErrorNotExistUser) {
		response := newResponse(ERROR, "El usuario o el password son incorrectos", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if data.ID <= 0 {
		response := newResponse(ERROR, "El usuario no existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := newResponse(ERROR, "Error al generar el token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", map[string]string{"token": token})
	return c.JSON(http.StatusOK, response)
}

// Update -> method of the class User (handler)
func (u *User) Update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(ERROR, "El id debe ser un número entero", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.User{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(ERROR, "La estructura no es la correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = u.storage.Update(ID, data)
	if errors.Is(err, storage.ErrorRowNull) {
		response := newResponse(ERROR, "No se aceptan datos nulos", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		response := newResponse(ERROR, "No se pudo crear el nuevo usuario", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", nil)
	return c.JSON(http.StatusOK, response)
}

// Delete -> method of the class User (handler)
func (u *User) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(ERROR, "El id debe ser un número entero", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = u.storage.Delete(ID)
	if errors.Is(err, storage.ErrorRowAffected) {
		response := newResponse(ERROR, "No se encontro el usuario", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(MESSAGE, "OK", nil)
	return c.JSON(http.StatusOK, response)

}
