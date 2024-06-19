package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akhil/to-do/service"

	"github.com/akhil/to-do/model"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createTodo model.TodoTask
	err := json.NewDecoder(r.Body).Decode(&createTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("\ntodo: %+v\n", createTodo)
	Id, err := service.CreateTodo(createTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successMsg := fmt.Sprintf("Your task is created with ID: %v", Id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMsg))
}
