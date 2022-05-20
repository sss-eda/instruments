package main

import (
	"log"
	"os"

	"github.com/sss-eda/instruments/internal/application/api"
)

func main() {
	envServerType, defined := os.LookupEnv("INSTRUMENTS_API_SERVER_TYPE")
	if !defined {
		envServerType = "OpenAPI"
	}

	var serverType api.ServerType
	err := serverType.UnmarshalText([]byte(envServerType))
	if err != nil {
		log.Fatal(err)
	}

	envGatewayType, defined := os.LookupEnv("INSTRUMENTS_API_GATEWAY_TYPE")
	if !defined {
		envGatewayType = "Memory"
	}

	var gatewayType api.GatewayType
	err = gatewayType.UnmarshalText([]byte(envGatewayType))
	if err != nil {
		log.Fatal(err)
	}

	app, err := api.NewApplication(serverType, gatewayType)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}
