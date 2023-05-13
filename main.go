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

	//u := &models.User{}
	//u.Name = "name"
	//u.Email = "test2@test.com"
	//u.PassWord = "password"
	//u.CreateUser()
	//
	//user, _ := models.GetUser(2)
	//user.CreateTodo("this is content.")

	//t, _ := models.GetTodo(1)
	//fmt.Println(t)

	//user, _ := models.GetUser(2)
	//user.CreateTodo("this is second content.")

	//todos, _ := models.GetTodos()
	//fmt.Println(todos)

	//user, _ := models.GetUser(2)

	t, _ := models.GetTodo(1)
	//t.Content = "updated2"
	//t.UserID = 3
	//t.UpdateTodo()

	t.DeleteTodo()

}
