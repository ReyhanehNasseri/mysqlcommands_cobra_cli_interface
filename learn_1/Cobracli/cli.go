package main

import (
	"fmt"
	dbmanagement "learn_1/DBmanagement"
	"log"
)

func main() {
	err := dbmanagement.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("successful connection")
	sqlcommands.AddCommand(insertCmd)
	sqlcommands.AddCommand(Updatecmd)
	sqlcommands.AddCommand(Deletecmd)
	sqlcommands.AddCommand(Readcmd)
	sqlcommands.AddCommand(Readallcmd)
	Execute()
}
