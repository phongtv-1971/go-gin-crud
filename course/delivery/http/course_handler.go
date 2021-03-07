package http

import (
	"go-gin-crud/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CourseHandler ...
type CourseHandler struct {
	CUsecase domain.CourseUseCase
}

// ResponseError ...
type ResponseError struct {
	Message string `json:"message"`
}

// NewCourseHandler ...
func NewCourseHandler(g *gin.Engine, us domain.CourseUseCase) {
	handler := &CourseHandler{
		CUsecase: us,
	}

	g.GET("/courses", handler.FetchCourse)
	g.POST("/courses", handler.Store)
	g.GET("/course/:id", handler.GetByID)
	g.PUT("/course/:id", handler.Update)
	g.DELETE("/course/:id", handler.Delete)
}

// FetchCourse return all course
func (c *CourseHandler) FetchCourse(g *gin.Context) {
	listC, err := c.CUsecase.Fetch()

	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	g.JSON(http.StatusOK, listC)
}

// Store create new record of course
func (c *CourseHandler) Store(g *gin.Context) {
	var course domain.Course
	err := g.Bind(&course)
	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = c.CUsecase.Store(&course)
	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	g.JSON(http.StatusCreated, course)
}

// GetByID return a course
func (c *CourseHandler) GetByID(g *gin.Context) {
	idP, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	id := int64(idP)
	res, err := c.CUsecase.GetByID(id)
	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	g.JSON(http.StatusCreated, res)
}

// Update return a course
func (c *CourseHandler) Update(g *gin.Context) {
	var course domain.Course
	err := g.Bind(&course)
	if err != nil {
		g.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	idP, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	id := int64(idP)
	err = c.CUsecase.Update(&course, id)

	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	g.JSON(http.StatusOK, course)
}

func (c *CourseHandler) Delete(g *gin.Context) {
	idP, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
		return
	}

	id := int64(idP)
	err = c.CUsecase.Delete(id)
	if err != nil {
		g.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	g.JSON(http.StatusNoContent, nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
