package helpers

import (
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	//"strings"
	"fmt"
	"os"
)

//type Config struct {
var authToken *jwtauth.JWTAuth
//}

func MakeToken(userId uuid.UUID) (string, error) {
	if err := godotenv.Load(".env"); err != nil{
		fmt.Println(err)
	}
	secret :=  os.Getenv("JWT_SECRET")
	authToken = jwtauth.New("HS256", []byte(secret), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, err := authToken.Encode(map[string]interface{}{"name": userId})
	if err != nil {
		return "", err
	}
	//fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenString, nil
}

func TokenValid(c *gin.Context) error {
	//tokenString := ExtractToken(c)
	//_, err :=

	//resp := jwtauth.Verify(authToken)

	

	return nil
}

//func ExtractToken(c *gin.Context) string {
//	token := c.Query("token")
//	if token != "" {
//		return token
//	}
//	bearerToken := c.Request.Header.Get("Authorization")
//	if len(strings.Split(bearerToken, " ")) == 2 {
//		return strings.Split(bearerToken, " ")[1]
//	}
//	return ""
//}

//func ExtractTokenID(c *gin.Context) (uint, error) {
//
//	tokenString := ExtractToken(c)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("API_SECRET")), nil
//	})
//	if err != nil {
//		return 0, err
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
//		if err != nil {
//			return 0, err
//		}
//		return uint(uid), nil
//	}
//	return 0, nil
//}