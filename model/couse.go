package model

// Course is -> (tb_courses) of the database.
type Course struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Active  bool   `json:"active"`
	Teacher string `json:"teacher"`
}
