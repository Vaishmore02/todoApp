package service

import "github.com/akhil/to-do/dao"

func DeleteTodo(Id string) error {
	return dao.DeleteTask(Id)
}
