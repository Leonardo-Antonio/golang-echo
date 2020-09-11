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
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "La estructura no es la correcta", nil, i)
		return response.JSONXML(c, response)
	}

	err = u.storage.SignIn(data)
	if errors.Is(err, storage.ErrorRowNull) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "No se aceptan datos nulos", nil, i)
		return response.JSONXML(c, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "No se pudo crear el nuevo usuario", nil, i)
		return response.JSONXML(c, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(c, response)
	}
	i := info{Code: http.StatusOK, Path: c.Path(), Method: "POST"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return response.JSONXML(c, response)
}

// LogIn -> method of the class User (handler)
func (u *User) LogIn(c echo.Context) error {
	data := model.User{}
	err := c.Bind(&data)
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "La estructura no es la correcta", nil, i)
		return response.JSONXML(c, response)
	}

	data, err = u.storage.LogIn(data)
	if errors.Is(err, storage.ErrorNotExistUser) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "El usuario o el password son incorrectos", nil, i)
		return response.JSONXML(c, response)
	}
	if data.ID <= 0 {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "El usuario no existe", nil, i)
		return response.JSONXML(c, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(c, response)
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: c.Path(), Method: "POST"}
		response := newResponse(ERROR, "Error al generar el token", nil, i)
		return response.JSONXML(c, response)
	}
	i := info{Code: http.StatusOK, Path: c.Path(), Method: "POST"}
	t := map[string]string{"token": token}
	response := newResponse(MESSAGE, "OK", t, i)
	return response.JSONXML(c, response)
}

// Update -> method of the class User (handler)
func (u *User) Update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "PUT"}
		response := newResponse(ERROR, "El id debe ser un número entero", nil, i)
		return response.JSONXML(c, response)
	}

	data := model.User{}
	err = c.Bind(&data)
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "PUT"}
		response := newResponse(ERROR, "La estructura no es la correcta", nil, i)
		return response.JSONXML(c, response)
	}

	err = u.storage.Update(ID, data)
	if errors.Is(err, storage.ErrorRowNull) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "PUT"}
		response := newResponse(ERROR, "No se aceptan datos nulos", nil, i)
		return response.JSONXML(c, response)
	}
	if errors.Is(err, storage.ErrorRowAffected) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "PUT"}
		response := newResponse(ERROR, "El usuario no existe", nil, i)
		return response.JSONXML(c, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: c.Path(), Method: "PUT"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(c, response)
	}

	i := info{Code: http.StatusOK, Path: c.Path(), Method: "PUT"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return response.JSONXML(c, response)
}

// Delete -> method of the class User (handler)
func (u *User) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "El id debe ser un número entero", nil, i)
		return response.JSONXML(c, response)
	}
	err = u.storage.Delete(ID)
	if errors.Is(err, storage.ErrorRowAffected) {
		i := info{Code: http.StatusBadRequest, Path: c.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "No se encontro el usuario", nil, i)
		return response.JSONXML(c, response)
	}
	if err != nil {
		i := info{Code: http.StatusInternalServerError, Path: c.Path(), Method: "DELETE"}
		response := newResponse(ERROR, "Ha ocurrido un error en la bd", nil, i)
		return response.JSONXML(c, response)
	}

	i := info{Code: http.StatusOK, Path: c.Path(), Method: "DELETE"}
	response := newResponse(MESSAGE, "OK", nil, i)
	return response.JSONXML(c, response)
}
