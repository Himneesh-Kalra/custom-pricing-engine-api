package main

import (
	"log"
	"net/http"

	"github.com/Himneesh-Kalra/custom-pricing-engine-api/router"
)

func main() {
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()).Error())
}
