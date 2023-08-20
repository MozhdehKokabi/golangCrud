package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var port string = os.Getenv("POSTGRES_PORT")
var host string = os.Getenv("POSTGRES_HOST")
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DBNAME")

var Db *sql.DB

func InitDataBase() error {
	portint, _ := strconv.ParseInt(port, 10, 64)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portint, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected!")

	Db = db
	fmt.Print(Db)

	return nil
}
