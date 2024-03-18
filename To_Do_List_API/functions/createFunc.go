package functions

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"to-do-api/strucks"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task strucks.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", "postgresql://root:root@db:5432/todolist?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (title, description, due_date, completed) VALUES ($1, $2, $3, $4)", task.Title, task.Description, task.DueDate, task.Completed)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}
