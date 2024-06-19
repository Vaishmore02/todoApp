package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akhil/to-do/service"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("Id")

	if id == "" {
		http.Error(w, "Id is missing", http.StatusBadRequest)
	}

	task, err := service.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
