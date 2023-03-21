package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ntekim/Altschool-Management-System/routes"
	"github.com/ntekim/Altschool-Management-System/config"
	"fmt"
	"log"
)

var Router routes.RouterConfig

func main() {
	_, err := config.ConnectToMySQL()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to MySQL Server!")

	//defer db.Close()
	
	router := Router.Routes()

	routerErr := router.Run("localhost:8088")
	if routerErr != nil{
		fmt.Println("Error", routerErr)
		panic(routerErr)
	}
	
	
}

//var DB *gorm.DB
//
//func ConnectToMySQL() (*gorm.DB, error){
//	err := godotenv.Load(".env")
//	if err != nil {
//		fmt.Println(err)
//	}
//	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/altschool", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"))
//
//	db, err := gorm.Open("mysql", dataSourceName)
//	if err != nil {
//		fmt.Println("err validating sql.open argument")
//		return nil, err
//	}
//
//	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Course{}, &models.CourseStudent{})
//	return db, nil
//}
//
//func GetDB() *gorm.DB {
//	return DB
//}
