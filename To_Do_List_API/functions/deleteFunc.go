package functions

import (
	"database/sql"
	"log"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Görev ID'si belirtilmedi", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", "postgresql://root:root@db:5432/todolist?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
