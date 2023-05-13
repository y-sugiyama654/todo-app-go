package controllers

import (
	"net/http"
	"todo-go-app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
