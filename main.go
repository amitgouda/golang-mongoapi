package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amitgouda/mongoapi/router"
)

func main() {
	fmt.Println("main")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))

}
