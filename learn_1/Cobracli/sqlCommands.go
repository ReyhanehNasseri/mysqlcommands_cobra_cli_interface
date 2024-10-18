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
		if err != nil {
			log.Println("Error converting age to integer:", err)
			return
		}

		newuser := dbmanagement.User{}
		newuser.ID = id
		newuser.Firstname = firstname
		newuser.Lastname = lastname

		saveduser, err := dbmanagement.InsertUser(newuser)
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
		id, _ := strconv.Atoi(args[0])
		err := dbmanagement.DeleteUser(id)
		if err != nil {
			log.Fatalf("Error deleting user: %v", err)
		}
		fmt.Printf("deleted user : %v\n", id)
	},
}

var Readallcmd = &cobra.Command{
	Use:   "Read_all",
	Short: "Read all users",
	Run: func(cmd *cobra.Command, args []string) {
		dbmanagement.Read_all()
	},
}
var Readcmd = &cobra.Command{
	Use:   "Read [id] ",
	Short: "Read a created user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		dbmanagement.Read(id)
	},
}

func Execute() {
	if err := sqlcommands.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
