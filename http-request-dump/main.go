package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
