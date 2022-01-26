package format

import (
	"fmt"
	"github.com/armanithiago/matrix-api/components/invert"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
)

// Invert receives a request and returns the matrix as a string in matrix format where the columns and rows are inverted
func Invert(w http.ResponseWriter, r *http.Request) {
	records, err := handlers.GetCsvFileFromRequest(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	fmt.Fprint(w, invert.Execute(records))
}
