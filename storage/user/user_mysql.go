package user

import (
	"database/sql"

	"github.com/Leonardo-Antonio/golang-echo/storage"

	"github.com/Leonardo-Antonio/golang-echo/model"
)

const (
	userCreate = "INSERT INTO tb_users VALUES (NULL, ?, ?)"
	userLogin  = "SELECT id ,email FROM tb_users WHERE email = ? AND pass = ?"
	userUpdate = "UPDATE tb_users SET email = ?, pass = ? WHERE id = ?"
	userDelete = "DELETE FROM tb_users WHERE id = ?"
)

// User -> class(attributes)
type User struct {
	db *sql.DB
}

// New -> method constructor
func New(db *sql.DB) *User {
	return &User{db}
}

// SignIn -> new user(create)
func (u *User) SignIn(user model.User) error {
	stmt, err := u.db.Prepare(userCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		storage.StringNull(user.Email),
		storage.StringNull(user.Password),
	)
	if err != nil {
		return storage.ErrorRowNull
	}

	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rA != 1 {
		return storage.ErrorRowAffected
	}
	return nil
}

// LogIn -> verify user
func (u *User) LogIn(user model.User) (data model.User, err error) {
	stmt, err := u.db.Prepare(userLogin)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Email, user.Password).Scan(
		&data.ID,
		&data.Email,
	)
	if err != nil {
		return data, storage.ErrorNotExistUser
	}
	return
}

// Update .
func (u *User) Update(id int, user model.User) error {
	stmt, err := u.db.Prepare(userUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		storage.StringNull(user.Email),
		storage.StringNull(user.Password),
		id,
	)
	if err != nil {
		return storage.ErrorRowNull
	}
	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rA != 1 {
		return storage.ErrorRowAffected
	}
	return nil
}

// Delete -> class of user
func (u *User) Delete(id int) error {
	stmt, err := u.db.Prepare(userDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rA, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rA != 1 {
		return storage.ErrorRowAffected
	}

	return nil
}
