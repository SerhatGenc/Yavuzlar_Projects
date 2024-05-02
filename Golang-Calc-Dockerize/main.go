package main

import (
	fnc "Calc/Functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calulator", fnc.HandleGo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
