package dbmanagement

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type User struct {
	ID        int `gorm:"primaryKey"`
	Firstname string
	Lastname  string
}

func ConnectDB() error {
	pass := os.Getenv("mysqlpassword")
	dsn := "root:" + pass + "@tcp(localhost:3306)/mydbtest"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	InitEntities()
	return nil
}

func InitEntities() {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalln("failed to create entity: ", err)
	}
}

func InsertUser(user User) (*User, error) {
	fmt.Println(user)
	result := db.Create(&user)

	return &user, result.Error
}

func UpdateUser(id int, newFirstname string, newLastname string) *User {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		log.Println(err)
	}

	user.Firstname = newFirstname
	user.Lastname = newLastname

	if err := db.Save(&user).Error; err != nil {
		log.Println("failed to update user: ", err)
	} else {
		log.Println("User updated successfully!")

	}
	return &user
}

func DeleteUser(id int) error {
	result := db.Delete(&User{}, id)
	return result.Error
}

func Read(id int) {

	var user []User
	db.First(&user, id)
	for _, attribute := range user {
		fmt.Printf("ID: %d, Firstname: %s , Lastname :%s\n", attribute.ID, attribute.Firstname, attribute.Lastname)
	}
}

func Read_all() {

	var Users []User
	db.Find(&Users)
	for _, user := range Users {
		fmt.Printf("ID: %d, Firstname: %s , Lastname : %s\n ", user.ID, user.Firstname, user.Lastname)
	}
}
