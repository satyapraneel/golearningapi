package main

import (
	"fmt"
	"log"
	"myapp/routes"
	"net/http"
)

func main() {
	fmt.Println("Server starting")

	err := http.ListenAndServe(":4000", routes.Router())
	if err != nil {
		log.Fatal(err)
	}

}
