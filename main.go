package main

import (
	"fmt"
	"go-rest-api-services/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
