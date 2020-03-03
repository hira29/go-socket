package main

import(
	"fmt"
	"go-socket/setup"
	"log"
	"net/http"
)

func main() {
	fmt.Println("WebSockets")
	setup.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8888", nil))
}