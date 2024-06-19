package dao

import (
	"fmt"
	"time"

	"github.com/akhil/to-do/model"
)

func AddTask(task model.TodoTask) error {
	query := `INSERT INTO dev.task (ID, UserID, Title, Description, Status, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Bind values to the query
	err := DB.session.Query(query, task.Id, task.UserId, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt).Exec()
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}

func DeleteTask(Id string) error {
	query := `DELETE FROM dev.task WHERE id = ?`

	err := DB.session.Query(query, Id).Exec()
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}

func GetTask(Id string) (model.TodoTask, error) {
	query := `SELECT ID, UserID, Title, Description, Status, CreatedAt, UpdatedAt FROM dev.task WHERE ID = ?`

	// Bind values to the query
	var data model.TodoTask
	err := DB.session.Query(query, Id).Scan(&data.Id, &data.UserId, &data.Title, &data.Description, &data.Status, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, fmt.Errorf("error executing query: %w", err)
	}

	return data, nil
}

func ListTask(UserID string) ([]*model.TodoTask, error) {

	// Example select statement to fetch multiple rows
	query := `SELECT ID, UserID, Title, Description, Status, CreatedAt, UpdatedAt FROM dev.task WHERE UserID = ? ALLOW FILTERING`

	// Execute the query
	iter := DB.session.Query(query, UserID).Iter()

	// Initialize slice to store fetched data
	var rows []*model.TodoTask

	// Iterate over the rows

	var Id string
	var UserId string
	var Title string
	var Description string
	var Status string
	var CreatedAt time.Time
	var UpdatedAt time.Time

	for iter.Scan(&Id, &UserId, &Title, &Description, &Status, &CreatedAt, &UpdatedAt) {
		// Append each row to the slice
		rows = append(rows, &model.TodoTask{
			Id:          Id,
			UserId:      UserId,
			Title:       Title,
			Description: Description,
			Status:      model.Status(Status),
			CreatedAt:   CreatedAt,
			UpdatedAt:   UpdatedAt,
		})
	}

	// Check for errors during iteration
	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("error closing iterator: %w", err)
	}

	return rows, nil
}
