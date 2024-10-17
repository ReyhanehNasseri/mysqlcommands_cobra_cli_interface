package main

import (
	"fmt"
	dbmanagement "learn_1/DBmanagement"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var sqlcommands = &cobra.Command{
	Use:   "mysql-template",
	Short: "My workbench with GORM and Cobra",
}
var insertCmd = &cobra.Command{
	Use:   "insert [id] [firstname] [lastname]",
	Short: "Insert a new user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		firstname := args[1]
		lastname := args[2]
		id, err := strconv.Atoi(args[0])
		fmt.Println(id)
		fmt.Println(firstname)
		fmt.Println(lastname)
		if err != nil {
			log.Println("Error converting age to integer:", err)
			return
		}

		newuser := dbmanagement.User{}
		newuser.ID = id
		newuser.Firstname = firstname
		newuser.Lastname = lastname
		fmt.Println(newuser.ID)
		fmt.Println(newuser.Firstname)
		fmt.Println(newuser.Lastname)

		saveduser, err := dbmanagement.InsertUser(newuser)
		fmt.Println(saveduser)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
		fmt.Printf("Inserted user : %v\n", saveduser)
	},
}

var Updatecmd = &cobra.Command{
	Use:   "update [id] [newfirstname] [newlastname]",
	Short: "update a created user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		newfirstname := args[1]
		newlastname := args[2]
		id, err := strconv.Atoi(args[0])
		updateduser := dbmanagement.UpdateUser(id, newfirstname, newlastname)
		fmt.Println(updateduser)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
		fmt.Printf("Updated user : %v\n", updateduser)
	},
}

var Deletecmd = &cobra.Command{
	Use:   "delete [id] ",
	Short: "delete a created user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		err1 := dbmanagement.DeleteUser(id)
		if err != nil {
			log.Fatalf("Error deleting user: %v", err1)
		}
		fmt.Printf("deleted user : %v\n", id)
	},
}

func Execute() {
	if err := sqlcommands.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
