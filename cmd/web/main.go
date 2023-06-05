package main

import (
	"fmt"
	"net/http"

	"github.com/cloudy4next/hotel_reservation_go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application port on: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
