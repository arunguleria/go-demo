package main

import (
	"fmt"
	"log"
	"net/http"

	"todo/todo-api/controller"
	"todo/todo-api/model"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
