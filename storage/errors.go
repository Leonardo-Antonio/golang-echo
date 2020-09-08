package storage

import (
	"errors"
)

var (
	// ErrorRowAffected .
	ErrorRowAffected = errors.New("No se afectaron las filas deseadas")
	// ErrorRowNull .
	ErrorRowNull = errors.New("No se puede ingresar valores nulos")
	// ErrorGetRowNull .
	ErrorGetRowNull = errors.New("No se puede obtener valores nulos")
	// ErrorNotExistUser .
	ErrorNotExistUser = errors.New("El usuario no existe")
	// ErrorNotExistCourse .
	ErrorNotExistCourse = errors.New("El curso no existe")
)
