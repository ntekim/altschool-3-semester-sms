package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ntekim/Altschool-Management-System/models"
	"net/http"
	"time"
	"fmt"
)

func (c *Config) GetStudentDetails(g *gin.Context) {
	id := g.Param("id")

	user, err := c.UserModel.GetUserDetail(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: fmt.Sprintf("User with ID: %s not found!", id),
				Data: err,
			},
		})
	}
	
	g.JSON(http.StatusOK, gin.H{
		"error": ResponsePayload{
			Error: false,
			Message: "Success",
			Data: user,
		},
	})
}

func (c *Config) GetUserCourses(g *gin.Context)  {
	id := g.Param("user_id")
	
	//Check if user_id exist before searching for courses
	_, err := c.UserModel.VerifyUser("id", id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": ResponsePayload{
				Error: true,
				Message: fmt.Sprintf("ID: %s not a valid user's id", id),
				Data: err,
			},
		})
	}

	var course models.Course

	c.CourseModel = &course

	courses, err := c.CourseModel.GetUserCourses("user_id", id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Error getting courses associated to user...",
				Data: err,
			},
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

func (c *Config) UpdateStudentCourseHandler(g *gin.Context)  {
	id := g.Param("id")
	var coursePayload *UpdateCourseRequestPayload

	if err := g.ShouldBindJSON(coursePayload); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var course *models.Course
	course.StudentID = coursePayload.StudentID 
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