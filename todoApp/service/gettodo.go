package service

import (
	"github.com/akhil/to-do/dao"
	"github.com/akhil/to-do/model"
)

func GetTodo(Id string) (model.TodoTask, error) {
	return dao.GetTask(Id)
}
