package math

import (
	"fmt"
	"github.com/armanithiago/matrix-api/components/sum"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
)

// Sum receives a request and returns the sum of the integers in the matrix
func Sum(w http.ResponseWriter, r *http.Request) {
	records, err := handlers.GetCsvFileFromRequest(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	fmt.Fprint(w, sum.Execute(records))
}
