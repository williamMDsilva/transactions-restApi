package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williamMDsilva/transactions-restApi/handles"
)

func RoutesApis() *mux.Router {

	routes := mux.NewRouter()

	routes.HandleFunc("/health", handles.HealthChecksHandler).Methods(http.MethodGet)
	routes.HandleFunc("/transaction", handles.CreateTransactionHandler).Methods(http.MethodPost)

	return routes
}
