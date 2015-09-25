package flow

import (
	"log"
	"net/http"
)

type visitHandler struct {
	visitted bool

	endpoint string
	method   string
}

func newVisitHandler(endpoint, method string) *visitHandler {
	return &visitHandler{endpoint: endpoint, method: method}
}

func (v *visitHandler) handle(response http.ResponseWriter, request *http.Request) {
	if v.method == request.Method {
		log.Printf("Visiting endpoint %v with HTTP method %v\n", v.endpoint, v.method)
		v.visitted = true

		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusNotFound)
}
