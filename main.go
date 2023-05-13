package main

import (
	"fmt"
	"log"
	"todo-go-app/app/controllers"
	"todo-go-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	err := controllers.StartMainServer()
	if err != nil {
		log.Fatalln(err)
	}
}
