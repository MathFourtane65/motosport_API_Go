package main

import (
	"api_go_mathieu_fourtane/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := route.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
