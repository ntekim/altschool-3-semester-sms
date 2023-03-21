package controllers

import (
	//"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/ntekim/Altschool-Management-System/helpers"
	"github.com/google/uuid"
	"github.com/ntekim/Altschool-Management-System/models"
)

type Config struct {
	UserModel *models.User
	ProfileModel *models.Profile
	CourseModel *models.Course
}

type RegisterRequestPayload struct {
	Firstname 	string 		`form:"firstname" json:"firstname" binding:"required,min=3"`
	Lastname  	string 		`form:"lastname" json:"lastname"  binding:"required,min=3"`
	Email     	string 		`form:"email" json:"email"  binding:"required,email"`
	Phone     	string 		`form:"phone" json:"phone" binding:"required"`
	Password    string		`form:"password" json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email     	string 		`form:"email" json:"email" binding:"required,email"`
	Password    string		`form:"password" json:"password" binding:"required,min=6"`
}

type UserResponsePayload struct {
	ID 			uuid.UUID		`json:"id"`
	Firstname	string			`json:"firstname"`
	Lastname	string			`json:"lastname"`
	Email 		string			`json:"email"`
	Phone 		string			`json:"phone"`
	Profile 	*models.Profile	`json:"profile"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
}

type ResponsePayload struct {
	Error	bool	`json:"error"`
	Message	string	`json:"message"`
	Data 	any 	`json:"data"`
}

func (c *Config) SignUpHandler(g *gin.Context)  {
	var userRequest RegisterRequestPayload

	routeParam := g.FullPath()
	
	if err := g.ShouldBindJSON(&userRequest); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	
	var user models.User
	user.ID = uuid.New()
	user.Firstname = userRequest.Firstname
	user.Lastname = userRequest.Lastname
	user.Email = userRequest.Email
	user.Phone = userRequest.Phone
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Error encountered while hashing the Password!",
			},
		})
	}
	user.Password = string(hashedPassword)
	switch routeParam {
	case "/create-user-account":
		user.Role = "user"
	case "/create-instructor-account":
		user.Role = "instructor"
	default:
		user.Role = "user"
	}
	user.CreatedAt = time.Now()

	var profile models.Profile
	model := Config{
		UserModel: &user,
		ProfileModel: &profile,
	}

	_, err = model.UserModel.VerifyUser("email", userRequest.Email)
	if err == nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "User with email already exists",
				Data: err,
			},
		})
		return
	}
	
	resp, err := model.UserModel.SignUp(&user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Error saving user's data!",
			},
		})
	}
	
	profile.ID = uuid.New()
	profile.UserID = resp.ID
	profile.CreatedAt = time.Now()
	
	profileResp, err := model.ProfileModel.CreateProfile(&profile)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Error saving user's data!",
			},
		})
	}
	
	var responsePayload UserResponsePayload
	
	responsePayload.ID	= user.ID
	responsePayload.Firstname = user.Firstname
	responsePayload.Lastname = user.Lastname
	responsePayload.Email = user.Email
	responsePayload.Phone = user.Phone
	responsePayload.Profile = profileResp
	responsePayload.CreatedAt = user.CreatedAt
	responsePayload.UpdatedAt = user.UpdatedAt

	g.JSON(http.StatusOK, gin.H{
		"response": ResponsePayload{
			Error: false,
			Message: "SignUp Successful",
			Data: responsePayload,
		},
	})
	
}

func (c *Config) LoginHandler(g *gin.Context)  {
	var user models.User
	var request LoginRequest

	config := Config{
		UserModel: &user,
	}
	
	if err := g.ShouldBindJSON(&request); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//fmt.Println(request.Email)
	userData, err := config.UserModel.VerifyUser("email", request.Email)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "User with email does not exist",
				Data: err,
			},
		})
		return
	}
	
	match, err := config.UserModel.PasswordMatches(request.Password, &userData)
	if match != true{
		g.JSON(http.StatusBadRequest, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Password does not match!",
				Data: err,
			},
		})	
	}

	token, err := helpers.MakeToken(userData.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"response": ResponsePayload{
				Error: true,
				Message: "Error occured while generating JWT token",
				Data: err,
			},
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"response": ResponsePayload{
			Error: false,
			Message: "Login Succesful",
			Data:	token,
		},
	})
}