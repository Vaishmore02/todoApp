package service

import (
	"github.com/akhil/to-do/dao"
	"github.com/akhil/to-do/model"
)

func ListTodo(UserId string) ([]*model.TodoTask, error) {
	return dao.ListTask(UserId)
}
