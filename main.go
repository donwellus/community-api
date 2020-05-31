package main

import (
	"community-api/schemas"
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	schema, _ := schemas.New()

	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	}))
	http.HandleFunc("/liveness", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
