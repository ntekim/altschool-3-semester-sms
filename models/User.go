package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"time"
	//"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type User struct {
	ID 		  	uuid.UUID	`json:"id" gorm:"primaryKey;autoIncrement:false"`
	Firstname 	string 		`json:"firstname"`
	Lastname  	string 		`json:"lastname"`
	Email     	string 		`json:"email" gorm:"type:varchar(100);unique_index"`
	Password	string		`json:"password"`
	Phone     	string 		`json:"phone"`
	Role      	string 		`gorm:"default:'user'" json:"role"` //["student", "instructor", "admin"]
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

var (
	DB *gorm.DB
)

func NewUserModelConfig(db *gorm.DB) {
	DB = db
	return 
}

func (u *User) VerifyUser(searchField, searchString string) (User, error) {
	field := fmt.Sprintf("%s = '%s'", searchField, searchString)
	user := User{}
	err := DB.Find(&user, field).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) ResetPassword() {

}

func (u User) SignUp(user *User) (*User, error) {
	err := DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) PasswordMatches(plainText string, user *User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (p *User) GetUserDetail(id string) (*User, error)  {
	var user *User
	err := DB.Model(User{}).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) DeleteUser(id string) (error){
	err := DB.Model(User{}).Where("id = ?", id).Delete(u).Error
	if err != nil{
		return err
	}

	return nil
}

func (u *User) GetAllStudents() ([]User, error) {
	var users []User
	err := DB.Model(User{}).Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) UpdateUser(id string) (*User, error) {
	err := DB.Model(User{}).Where("id = ?", id).Update(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}