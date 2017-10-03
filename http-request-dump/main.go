package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	poweredBy := os.Getenv("POWERED_BY")
	if len(poweredBy) == 0 {
		poweredBy = "Go"
	}
	fmt.Fprintf(w, "Powered by %v on %v\n\n", poweredBy, hostname)

	r.Write(w)
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
