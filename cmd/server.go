package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
