package usecase

import (
	"go-gin-crud/domain"
)

type courceUsecase struct {
	courseRepo domain.CourseRepository
}

// NewCourseUsecase implement interface domain.CourseUseCase
func NewCourseUsecase(c domain.CourseRepository) domain.CourseUseCase {
	return &courceUsecase{
		courseRepo: c,
	}
}

func (c *courceUsecase) Fetch() (res []domain.Course, err error) {
	res, err = c.courseRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return
}

func (c *courceUsecase) GetByID(id int64) (res domain.Course, err error) {
	res, err = c.courseRepo.GetByID(id)
	return
}

func (c *courceUsecase) Update(co *domain.Course, id int64) (err error) {
	err = c.courseRepo.Update(co, id)
	return
}

func (c *courceUsecase) Store(co *domain.Course) (err error) {
	err = c.courseRepo.Store(co)
	return
}

func (c *courceUsecase) Delete(id int64) error {
	return nil
}
