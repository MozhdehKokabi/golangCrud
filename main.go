package main

import (
	"ARCAPTCHA/admin"
	"ARCAPTCHA/handlers"
	"ARCAPTCHA/repository"
	"fmt"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)


func main() {
	err := repository.InitDataBase()
	if err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	
	e.POST("/login", handlers.LogIn)
	e.DELETE("/", admin.DeleteUser)
	e.PUT("/", admin.UpdateUser)
	e.GET("/", admin.ReadUser)
	e.POST("/signup", handlers.SignUp)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))

}
