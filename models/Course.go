package models

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

type Course struct {
	ID         		uuid.UUID 	`json:"id"`
	Title      		string 		`json:"title"`
	Desc       		string 		`json:"desc,omitempty"`
	Code       		string 		`json:"course_code,omitempty"`
	Status	   		string		`json:"status" gorm:"default:'Not Enrolled'"`
	InstructorID	string		`json:"instructor_id"`
	StudentID 	   	string 		`json:"user_id"`
	CreatedAt  	   	time.Time 	`json:"created_at"`
	UpdatedAt  	   	time.Time 	`json:"updated_at"`
}

//type UserCourse struct {
//	ID        string   	`json:"id"`
//	CourseID  string 	`json:"course_id"`
//	UserID 	  string   	`json:"student_id"`
//	Status    string   	`json:"status"` //Wishlist || Cart || Enrolled
//	CreatedAt string   	`json:"created_at"`
//	UpdatedAt string   	`json:"updated_at"`
//}

func (c *Course) AddCourse(coursePayload *Course) (*Course, error){
	err := DB.Create(coursePayload).Error
	if err != nil{
		return nil, err
	}
	return coursePayload, nil
}

//func (c *Course) AddStudentCourseInfo(courseStudentPayload *UserCourse) (*UserCourse, error){
//	err := DB.Create(courseStudentPayload).Error
//	if err != nil{
//		return nil, err
//	}
//	
//	return courseStudentPayload, nil
//}

func (c *Course) GetCourse(id string) (Course, error) {
	var course Course
	err := DB.Model(Course{}).Where("id = ?", id).First(course).Error
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func (c *Course) GetUserCourses(searchField, searchString string) ([]Course, error) {
	var courses []Course
	query := fmt.Sprintf("%s = '%s'", searchField, searchString)

	err :=  DB.Model(Course{}).Where(query).Find(courses).Error
	if err != nil{
		return nil, err
	}
	return courses, nil
}

func (c *Course) GetAllCourse() ([]Course, error) {
	var courses []Course
	err := DB.Model(Course{}).Find(courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *Course) UpdateCourse(id string, coursePayload *Course) (*Course, error) {
	err := DB.Model(Course{}).Where("id = ?", id).Update(coursePayload).Error
	if err != nil {
		return nil, err
	}
	return coursePayload, nil
}

//func (c *Course) UpdateStudentCourseInfo(courseID, userID string, studentCoursePayload *CourseStudent) (*CourseStudent, error) {
//	err := DB.Model(CourseStudent{}).Where("course_id = ?", courseID).First(studentCoursePayload).Error
//	if err != nil{
//		return nil, err
//	}
//	
//	return studentCoursePayload, nil
//}

func (c *Course) DeleteCourse(id string) error {
	err := DB.Model(Course{}).Where("id = ?", id).Delete(&c).Error
	if err != nil{
		return err
	}
	return nil
}

//func (c *Course) DeleteStudentCourseInfo(id string) error {
//	var CourseStudentPayload CourseStudent
//	err := DB.Model(CourseStudent{}).Where("id = ?", id).Delete(&CourseStudentPayload).Error
//	if err != nil{
//		return err
//	}
//	return nil
//}
