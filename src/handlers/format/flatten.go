package format

import (
	"fmt"
	"github.com/armanithiago/matrix-api/components/flatten"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
)

// Flatten receives a request and returns the matrix as a 1 line string, with values separated by commas.
func Flatten(w http.ResponseWriter, r *http.Request) {
	records, err := handlers.GetCsvFileFromRequest(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	fmt.Fprint(w, flatten.Execute(records))
}
