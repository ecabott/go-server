package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./Content")))
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
