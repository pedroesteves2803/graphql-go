package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4	)",
		id, name, description, categoryID)

	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil

}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	course := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		course = append(course, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}

	return course, nil

}

func (c *Course) FindByCategoryID(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Courses := []Course{}
	for rows.Next() {
		var id, name, description, category string
		if err := rows.Scan(&id, &name, &description, &category); err != nil {
			return nil, err
		}
		Courses = append(Courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryId})
	}
	return Courses, nil
}
