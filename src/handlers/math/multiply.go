package math

import (
	"fmt"
	"github.com/armanithiago/matrix-api/components/multiply"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
)

// Multiply receives a request and returns the product of the integers in the matrix
func Multiply(w http.ResponseWriter, r *http.Request) {
	records, err := handlers.GetCsvFileFromRequest(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	fmt.Fprint(w, multiply.Execute(records))
}
