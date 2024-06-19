package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhil/to-do/dao"
	"github.com/akhil/to-do/handler"
	"github.com/gorilla/mux"
)

func main() {
	dao.InitDB()
	initRouter()
}
func initRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todo", handler.GetTodo).Methods(http.MethodGet)
	r.HandleFunc("/todo/list", handler.ListTodo).Methods(http.MethodGet)
	r.HandleFunc("/todo", handler.DeleteTodo).Methods(http.MethodDelete)

	fmt.Println("Server started on 8080 ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
