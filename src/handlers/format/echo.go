package format

import (
	"fmt"
	"github.com/armanithiago/matrix-api/components/echo"
	"github.com/armanithiago/matrix-api/handlers"
	"net/http"
)

// Echo receives a request and returns the matrix as a string in matrix format
func Echo(w http.ResponseWriter, r *http.Request) {
	records, err := handlers.GetCsvFileFromRequest(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}
	fmt.Fprint(w, echo.Execute(records))
}
