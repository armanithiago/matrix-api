package routes

import (
	"github.com/armanithiago/matrix-api/handlers/format"
	"github.com/armanithiago/matrix-api/handlers/math"
	"net/http"
)

// Configure setup all the route configuration
func Configure() {
	http.HandleFunc("/echo", format.Echo)
	http.HandleFunc("/sum", math.Sum)
}
