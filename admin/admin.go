package admin

import (
	"ARCAPTCHA/auth"
	"ARCAPTCHA/models"
	"ARCAPTCHA/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// ReadUser
// @Summary Get User Information
// @Accept  json
// @Produce  json
// @Success 204 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router / [GET]

func ReadUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		print(err)
		// return err
	}
	newPerson := models.ReqUser{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	token, err := auth.GenerateJWT(req.UserName)
	var jwttoken = fmt.Sprintf("token: " + token)
	fmt.Println(jwttoken)
	c.Response().Header().Set("Authorization", jwttoken)
	fmt.Println("successfully set")

	var password string
	var email string
	var phone int
	var address string
	err = repository.Db.QueryRow("select password, email, phone, address from students where username = $1", req.UserName).Scan(&password, &email, &phone, &address)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err == nil {
		newPerson.Address = address
		newPerson.Email = email
		newPerson.Phone = phone

		fmt.Println(newPerson)
		fmt.Println("Password was correct!")
		return c.JSON(http.StatusOK, newPerson)

	}

	return c.JSON(http.StatusOK, "Wrong UserName")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func UpdateUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		print(err)

	}
	newPerson := models.ReqUser{
		Password: req.Password,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
	}

	token, err := auth.GenerateJWT(req.UserName)
	var jwttoken = fmt.Sprintf("token: " + token)
	fmt.Println(jwttoken)
	c.Response().Header().Set("Authorization", jwttoken)
	fmt.Println("successfully set")

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	newPerson.Password = "123"

	// fmt.Println(newPerson)

	_ = repository.Db.QueryRow("UPDATE students SET password= $2, email= $3, phone= $4, address= $5 where username = $1 ", req.UserName, string(hash), req.Email, req.Phone, req.Address)
	// fmt.Println(e)

	return c.JSON(http.StatusOK, newPerson)

	// return c.JSON(http.StatusOK, "Wrong Username")
}

// DeleteUser
// @Summary delete resource admin
// @Accept  json
// @Produce  json
// @Success 204 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router / [delete]

func DeleteUser(c echo.Context) error {
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		print(err)
		// return err
	}

	var name string
	err = repository.Db.QueryRow("select username from students where username= $1", req.UserName).Scan(&name)

	if name == req.UserName {
		result := repository.Db.QueryRow("delete from students where username= $1", req.UserName)
		fmt.Println(result)

		return c.JSON(http.StatusOK, "Successfully Delete")
	}
	return c.JSON(http.StatusOK, "Wrong Username")
}
