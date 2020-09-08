package course

import (
	"database/sql"

	"github.com/Leonardo-Antonio/golang-echo/storage"

	"github.com/Leonardo-Antonio/golang-echo/model"
)

const (
	courseCreate  = "INSERT INTO tb_courses VALUES (NULL, ?, ?, ?)"
	courseGetAll  = "SELECT id, name, active, teacher FROM tb_courses"
	courseGetByID = "SELECT id, name, active, teacher FROM tb_courses WHERE id = ?"
	courseUpdate  = "UPDATE tb_courses SET name = ?, active = ?, teacher = ? where id = ?"
	courseDelete  = "DELETE FROM tb_courses where id = ?"
)

// Course -> class
type Course struct {
	db *sql.DB
}

// New -> method constructor
func New(db *sql.DB) *Course {
	return &Course{db}
}

// Create -> method of the class Course
func (c *Course) Create(course model.Course) error {
	stmt, err := c.db.Prepare(courseCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		storage.StringNull(course.Name),
		course.Active,
		storage.StringNull(course.Teacher),
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

// GetAll -> method of the class Course
func (c *Course) GetAll() (courses []model.Course, err error) {
	stmt, err := c.db.Prepare(courseGetAll)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}

	// valid data null
	teacherNull := sql.NullString{}

	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Active,
			&teacherNull,
		)
		if err != nil {
			return nil, storage.ErrorGetRowNull
		}
		course.Teacher = teacherNull.String
		courses = append(courses, course)
	}
	return
}

// GetByID -> method of the class Course
func (c *Course) GetByID(id int) (course model.Course, err error) {
	stmt, err := c.db.Prepare(courseGetByID)
	if err != nil {
		return
	}
	defer stmt.Close()

	//valid data null
	teacherNull := sql.NullString{}

	err = stmt.QueryRow(id).Scan(
		&course.ID,
		&course.Name,
		&course.Active,
		&teacherNull,
	)
	if err != nil {
		return course, storage.ErrorNotExistCourse
	}
	course.Teacher = teacherNull.String
	return
}

// Update -> method of the class Course
func (c *Course) Update(id int, course model.Course) error {
	stmt, err := c.db.Prepare(courseUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		course.Name,
		course.Active,
		storage.StringNull(course.Teacher),
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
		return storage.ErrorNotExistCourse
	}
	return nil
}

// Delete -> method of the class Course
func (c *Course) Delete(id int) error {
	stmt, err := c.db.Prepare(courseDelete)
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
		return storage.ErrorNotExistCourse
	}

	return nil
}
