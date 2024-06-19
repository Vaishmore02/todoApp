package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akhil/to-do/service"
)

func ListTodo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("UserId")

	if id == "" {
		http.Error(w, "UserId is missing", http.StatusBadRequest)
	}

	tasks, err := service.ListTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
