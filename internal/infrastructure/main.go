package main

import (
	"log"
	"net/http"

	"github.com/sss-eda/instruments/internal/infrastructure/graphql"
)

func main() {
	handler, err := graphql.NewHTTPHandler()
	if err != nil {
		log.Fatal(err)
	}

	// h := handler.New(&handler.Config{
	// 	Schema:   &schema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
