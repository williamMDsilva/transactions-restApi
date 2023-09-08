package handles

import (
	"io"
	"net/http"
)

// POST /transaction
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"message": "Operation finish with success","status": "SUCCESS"}`)
}
