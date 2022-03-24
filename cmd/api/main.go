package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sss-eda/instruments/internal/application/api"
)

func main() {
	envAPIType, defined := os.LookupEnv("INSTRUMENTS_API_TYPE")
	if !defined {
		envAPIType = "OpenAPI"
	}
	var apiType api.Type
	err := apiType.UnmarshalText([]byte(envAPIType))
	if err != nil {
		log.Fatal(err)
	}

	handler, err := api.NewHandler(apiType)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", handler))
}
