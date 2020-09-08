package storage

import "github.com/Leonardo-Antonio/golang-echo/model"

type (
	// Course -> interface implemented in handlers
	Course interface {
		Create(model.Course) error
		GetAll() ([]model.Course, error)
		GetByID(int) (model.Course, error)
		Update(int, model.Course) error
		Delete(int) error
	}

	// User -> interface implemented in handlers
	User interface {
		SignIn(model.User) error
		LogIn(model.User) (int, error)
		Update(int, model.User) error
		Delete(int) error
	}
)
