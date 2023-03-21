package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ntekim/Altschool-Management-System/controllers"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	//"net/http"
	"fmt"
	"os"
)

// @title Altschool SMS API Specification
// @version 1.0
// @description Student management System server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email admin@altschoom-sms.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

type RouterConfig struct {
	controllerConfig *controllers.Config
}
var authToken *jwtauth.JWTAuth

func init()  {
	if err := godotenv.Load(".env"); err != nil{
		fmt.Println(err)
	}
	authToken = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

func (app *RouterConfig) Routes() *gin.Engine{
	router := gin.Default()

	public := router.Group("/api")

	//public.Use(func(handler http.Handler) http.Handler {
	//	jwtauth.Verifier(authToken)
	//})
	public.POST("/create-user-account", app.controllerConfig.SignUpHandler)
	public.POST("/create-instructor-account", app.controllerConfig.SignUpHandler)
	public.POST("/login", app.controllerConfig.LoginHandler)

	protected := router.Group("/api")
	protected.POST("/add-course", app.controllerConfig.AddCourseHandler)
	protected.POST("/list-courses", app.controllerConfig.ListCourses)
	protected.POST("/user-courses/:user_id", app.controllerConfig.GetUserCourses)
	//protected.POST("/instructor-courses/:id", app.controllerConfig.GetUserCourses)
	protected.PUT("/update-course/:id", app.controllerConfig.UpdateCourseHandler)
	protected.PUT("/update-student-course/:id", app.controllerConfig.UpdateStudentCourseHandler)


	return router

}
