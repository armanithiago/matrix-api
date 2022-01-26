package main

import (
	"github.com/armanithiago/matrix-api/routes"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
func main() {
	routes.Configure()
	http.ListenAndServe(":8080", nil)
}
