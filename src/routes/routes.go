package routes

import (
	"github.com/armanithiago/matrix-api/handlers/format"
	"github.com/armanithiago/matrix-api/handlers/math"
	"net/http"
)

// Configure setup all the route configuration
func Configure() {
	http.HandleFunc("/echo", format.Echo)
	http.HandleFunc("/flatten", format.Flatten)
	http.HandleFunc("/invert", format.Invert)
	http.HandleFunc("/sum", math.Sum)
	http.HandleFunc("/multiply", math.Multiply)

}
