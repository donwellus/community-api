package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/liveness", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
