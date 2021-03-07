package mysql

import (
	"go-gin-crud/domain"

	"gorm.io/gorm"
)

type mysqlCourseRepository struct {
	Conn *gorm.DB
}

// NewMysqlCourseRepository return domain.CourseRepository interface
func NewMysqlCourseRepository(Conn *gorm.DB) domain.CourseRepository {
	return &mysqlCourseRepository{Conn}
}

func (m *mysqlCourseRepository) Fetch() ([]domain.Course, error) {
	res := []domain.Course{}
	err := m.Conn.Preload("CourseContents").Find(&res)

	return res, err.Error
}

func (m *mysqlCourseRepository) GetByID(id int64) (domain.Course, error) {
	res := domain.Course{}
	err := m.Conn.Where("id = ?", id).Preload("CourseContents").Find(&res)

	if res.ID == 0 {
		err.Error = domain.ErrNotFound
	}
	return res, err.Error
}

func (m *mysqlCourseRepository) Update(co *domain.Course, id int64) error {
	res := domain.Course{}
	query := m.Conn.Where("id = ?", id).Find(&res)

	if res.ID == 0 {
		query.Error = domain.ErrNotFound
	} else {
		query.Save(&co)
	}
	return query.Error
}

func (m *mysqlCourseRepository) Store(co *domain.Course) error {
	err := m.Conn.Create(&co)

	return err.Error
}

func (m *mysqlCourseRepository) Delete(id int64) error {
	return nil
}
