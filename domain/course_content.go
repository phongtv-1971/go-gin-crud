package domain

import "gorm.io/gorm"

// CourseContent ...
type CourseContent struct {
	Title       string
	Description string
	CourseID    uint
	course      Course
	gorm.Model
}

// CourseContentRepository ...
type CourseContentRepository interface {
	GetById(id int64) (CourseContent, error)
}
