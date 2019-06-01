package main

import (
	"Grocery-Shopping-User-Module/src/app/api"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func main() {
	fmt.Println("Starting User Module.... ")
	config := api.GetConfig()

	app := &api.App{}
	app.Initialize(config)

	port := os.Getenv("PORT")
	fmt.Println("Port No. is :"+port)
	app.Run(":"+port)

	fmt.Println("Started User Module.... ")
}