package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB
var SqlDB *sql.DB
var err error

type user struct {
	First_name string
	Last_name  string
	Email      string
	Phone      string
	Password   string
	Username   string
}

func AddUser(firstName string, lastname string, email string, phone string, password string, username string) bool {
	var user user
	user.First_name = firstName
	user.Last_name = lastname
	user.Email = email
	user.Phone = phone
	user.Password = password
	user.Username = username

	tx := DBConn.Create(&user)

	if tx.Error != nil {
		return false
	}
	return true
}

func FindUser(username string, lastName string, email string) bool {
	var user user
	tx := DBConn.Where("username = ? AND last_Name = ? AND email = ? ", username, lastName, email).Find(&user)
	if tx.Error != nil {
		return false
	}
	return true
}
func UserLogin(username string, password string) bool {

	var user user
	DBConn.Where("username = ? AND password = ?", username, password).Find(&user)
	if user.Username == username && username != "" {
		return true
	}
	return false
}
func ConnectDB() {
	//connecting to database
	dsn := "host=localhost user=postgres password=12344321 dbname=postgres port=5432 sslmode=disable"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	SqlDB, err = DBConn.DB()
	if err != nil {
		panic(err)
	}
	SqlDB.Ping()
	fmt.Println("Successfully connected database!")
}
