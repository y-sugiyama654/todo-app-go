package main

import (
	"fmt"
	"todo-go-app/app/models"
)

func main() {
	//fmt.Println(config.Config.Port)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.DbName)
	//fmt.Println(config.Config.LogFile)
	//
	//log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "name"
	u.Email = "test2@test.com"
	u.PassWord = "password"
	u.CreateUser()
}
