package handler

import (
	"fmt"
	"net/http"

	"github.com/akhil/to-do/service"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("Id")

	if id == "" {
		http.Error(w, "Id is missing", http.StatusBadRequest)
	}
	err := service.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	successMsg := fmt.Sprintf("Your task is deleted with ID: %v", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMsg))
}
