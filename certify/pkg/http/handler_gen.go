// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/certify/pkg/endpoint"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeGetAuthInfoHandler(m, endpoints, options["GetAuthInfo"])
	makePostAuthInfoHandler(m, endpoints, options["PostAuthInfo"])
	makeGetAllUnCertifyHandler(m, endpoints, options["GetAllUnCertify"])
	makeGetUnCertifyInfoHandler(m, endpoints, options["GetUnCertifyInfo"])
	makePostCertifyStateHandler(m, endpoints, options["PostCertifyState"])
	return m
}
