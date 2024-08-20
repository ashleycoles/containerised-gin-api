package main

import (
	"fmt"
)

func main() {
	router := router()

	if err := router.Run("0.0.0.0:8080"); err != nil {
		fmt.Println(err.Error())
	}
}
