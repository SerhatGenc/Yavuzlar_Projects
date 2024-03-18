package main

import (
	"fmt"
	"log"
	"net/http"
	"to-do-api/functions"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/tasks", functions.ListTasksHandler)
	http.HandleFunc("/tasks/create", functions.CreateTaskHandler)
	http.HandleFunc("/tasks/delete", functions.DeleteTaskHandler)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
