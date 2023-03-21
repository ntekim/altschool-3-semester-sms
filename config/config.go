package config

import (
	//"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ntekim/Altschool-Management-System/models"
	"fmt"
	"os"
)

func ConnectToMySQL() (*gorm.DB, error){
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//Dbdriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")
	//DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(mysql:3306)/altschool?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword)
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/altschool?parseTime=true", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"))

	db, err := gorm.Open("mysql", DBURL)
	if err != nil {
		fmt.Println("Cannot connect to mysql database ")
		return nil, err
	}
	
	db.AutoMigrate(&models.User{}, &models.Course{}, &models.Profile{})
	
	defer func() {
		models.NewUserModelConfig(db)
	}()
	
	return db, nil
}

//func GetDB() *gorm.DB {
//	return DB
//}