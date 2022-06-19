package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/huzaifamk/mongo-api/router"
)

func main() {
	fmt.Println("MongoDB API testing by Hmk")
	r := router.Router()
	fmt.Println("This server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening & Serving at port 4000 ...")
}
