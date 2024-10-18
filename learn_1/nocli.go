package main

import (
	"fmt"
	dbmanagement "learn_1/DBmanagement"
	"log"
)

func main() {
	var id int

	//Connect
	err := dbmanagement.ConnectDB()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("successful connection")

	//Insert
	newuser := dbmanagement.User{}
	fmt.Println("ID : ")
	fmt.Scan(&newuser.ID)
	fmt.Println("Firstname : ")
	fmt.Scan(&newuser.Firstname)
	fmt.Println("Lastname : ")
	fmt.Scan(&newuser.Lastname)
	saveduser, err := dbmanagement.InsertUser(newuser)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	fmt.Printf("Inserted user : %v\n", saveduser)

	//Update
	var newFirstName string
	var newLastname string
	fmt.Println("enter user ID to update:")
	fmt.Scan(&id)
	fmt.Println("Enter the new first name :")
	fmt.Scan(&newFirstName)
	fmt.Println("Enter the new last name :")
	fmt.Scan(&newLastname)
	editeduser := dbmanagement.UpdateUser(id, newFirstName, newLastname)
	fmt.Printf("Edited user : %v\n", editeduser)

	//Delete
	var DeleteUser_id int
	fmt.Println("enter id to delete the user : ")
	fmt.Scan(&DeleteUser_id)
	if err := dbmanagement.DeleteUser(DeleteUser_id); err != nil {
		log.Fatalf("Error deleting user : %v", err)
	}
	fmt.Printf("user %v deleted \n", DeleteUser_id)

	//Read
	fmt.Println("insert id to read: ")
	fmt.Scan(&id)
	dbmanagement.Read(id)

	//Read all users
	dbmanagement.Read_all()

}
