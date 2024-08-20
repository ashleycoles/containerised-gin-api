package main

import (
	"ashleycoles/logbook-api/routing"
	"ashleycoles/logbook-api/services"
	"fmt"
)

func main() {
	services.LoggerSetup()
	router := routing.Router()

	if err := router.Run("0.0.0.0:8080"); err != nil {
		fmt.Println(err.Error())
	}
}
