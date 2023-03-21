package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ntekim/Altschool-Management-System/models"
	"net/http"
	"github.com/google/uuid"
	"time"
	"fmt"
)
type CreateCourseRequestPayload struct {
	Title      		string `json:"title"`
	Desc       		string `json:"desc,omitempty"`
	Code       		string `json:"course_code,omitempty"`
	InstructorID	string
}

type UpdateCourseRequestPayload struct {
	Title      string `json:"title,omitempty"`
	Desc       string `json:"desc,omitempty"`
	Code       string `json:"course_code,omitempty"`
	Status	   string `json:"status,omitempty"`
	StudentID  string `json:"user_id,omitempty"`
}


func (c *Config) UpdateCourseHandler(g *gin.Context)  {
	id := g.Param("id")
	
	var coursePayload *UpdateCourseRequestPayload
	
	if err := g.ShouldBindJSON(coursePayload); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	
	var course *models.Course
	course.Title = coursePayload.Title
	course.Desc = coursePayload.Desc
	course.Code = coursePayload.Code
	course.Status = coursePayload.Status
	course.UpdatedAt = time.Now()
	
	c.CourseModel = course
	
	courseResp, err := c.CourseModel.UpdateCourse(id, course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Encountered an error while creating course",
				Data: err,
			},
		})
		return
	}
	
	g.JSON(http.StatusOK, gin.H{
		"response": ResponsePayload{
			Error: false,
			Message: "success",
			Data: courseResp,
		},
	})
}

func (c *Config) AddCourseHandler(g *gin.Context)  {
	var coursePayload *CreateCourseRequestPayload

	if err := g.ShouldBindJSON(coursePayload); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var course *models.Course
	course.ID = uuid.New()
	course.Title = coursePayload.Title
	course.Desc = coursePayload.Desc
	course.Code = coursePayload.Code
	course.InstructorID = coursePayload.InstructorID
	course.CreatedAt = time.Now()
	
	c.CourseModel = course

	courseResp, err := c.CourseModel.AddCourse(course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Encountered an error while creating course",
				Data: err,
			},
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"response": ResponsePayload{
			Error: false,
			Message: "success",
			Data: courseResp,
		},
	})
}

func (c *Config) ListCourses(g *gin.Context)  {
	
	c.CourseModel = &models.Course{}
	
	courses, err :=  c.CourseModel.GetAllCourse()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	g.JSON(http.StatusOK, gin.H{
		"response": ResponsePayload{
			Error: false,
			Message: "Success",
			Data: courses,
		},
	})
}

func (c *Config) GetCourseDetails(g *gin.Context) {
	id := g.Param("id")

	c.CourseModel = &models.Course{}
	course, err := c.CourseModel.GetCourse(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: fmt.Sprintf("Course with ID: %s not found!", id),
				Data: err,
			},
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"error": ResponsePayload{
			Error: false,
			Message: "Success",
			Data: course,
		},
	})
}

