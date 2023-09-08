// endpoints_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/williamMDsilva/transactions-restApi/handles"
)

func TestCreateTransactionHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/transaction", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handles.CreateTransactionHandler)

	handler.ServeHTTP(rr, req)

	expected := `{"message": "Operation finish with success","status": "SUCCESS"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
