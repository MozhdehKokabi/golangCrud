package handlers

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

// SignUp
// @Summary User Signup
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body models.ReqUser true "models.ReqUser"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /SignUp [post]
func SignUp(c echo.Context) error {

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

	fmt.Print(newPerson)

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(newPerson.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	newPerson.Password = string(hash)

	insertDynStmt := `insert into "students" ("password", "username", "email", "phone", "address", "role") values ($1, $2, $3, $4, $5, $6)`

	result, err := repository.Db.Exec(insertDynStmt, newPerson.Password, newPerson.UserName, newPerson.Email, newPerson.Phone, newPerson.Address, "User")

	fmt.Println(err)
	fmt.Println(result)

	return c.JSON(http.StatusOK, "Successfully SignUp")
}

// Login
// @Summary login handlers
// @Accept  json
// @Produce  json
// @Param    req  body models.ReqUser true "models.ReqUser"
// @Success 200 {object} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /Login [post]

func LogIn(c echo.Context) error {
	// Your login logic here
	var req models.ReqUser
	err := c.Bind(&req)
	if err != nil {
		print(err)
		// return err
	}
	token, err := auth.GenerateJWT(req.UserName)
	var jwttoken = fmt.Sprintf("token: " + token)
	fmt.Println(jwttoken)
	c.Response().Header().Set("Authorization", jwttoken)
	fmt.Println("successfully set")

	var password string
	err = repository.Db.QueryRow("select password from students where username= $1", req.UserName).Scan(&password)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err == nil {

		fmt.Println("Password was correct!")
		return c.JSON(http.StatusOK, "Successfully Login")
	}

	return c.JSON(http.StatusNonAuthoritativeInfo, "Wrong username or password")

}
