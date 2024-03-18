package functions

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"to-do-api/strucks"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", "postgresql://root:root@db:5432/todolist?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []strucks.Task

	for rows.Next() {
		var task strucks.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
