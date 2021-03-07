package domain

import "gorm.io/gorm"

// Course ...
type Course struct {
	Title          string
	CourseContents []CourseContent
	gorm.Model
}

// CourseUseCase ...
type CourseUseCase interface {
	Fetch() ([]Course, error)
	GetByID(id int64) (Course, error)
	Update(co *Course, id int64) error
	Store(co *Course) error
	Delete(id int64) error
}

// CourseRepository ...
type CourseRepository interface {
	Fetch() ([]Course, error)
	GetByID(id int64) (Course, error)
	Update(co *Course, id int64) error
	Store(co *Course) error
	Delete(id int64) error
}
