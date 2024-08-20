package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	log.SetOutput(file)

	router := router()

	if err := router.Run("0.0.0.0:8080"); err != nil {
		fmt.Println(err.Error())
	}
}
