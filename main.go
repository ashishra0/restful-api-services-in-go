package main

import (
	"fmt"
	"go-rest-api-services/handlers"
	"net/http"
	"os"
<<<<<<< HEAD
	"restful-api-services-in-go/handlers"
=======
>>>>>>> d1b0bf892e8c11042df951a924e39f304d2ff660
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
