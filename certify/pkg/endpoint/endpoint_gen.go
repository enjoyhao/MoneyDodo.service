// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/certify/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetAllUnAuthEndpoint    endpoint.Endpoint
	PostAuthInfoEndpoint    endpoint.Endpoint
	PostCertifyInfoEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.CertifyService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		GetAllUnAuthEndpoint:    MakeGetAllUnAuthEndpoint(s),
		PostAuthInfoEndpoint:    MakePostAuthInfoEndpoint(s),
		PostCertifyInfoEndpoint: MakePostCertifyInfoEndpoint(s),
	}
	for _, m := range mdw["GetAllUnAuth"] {
		eps.GetAllUnAuthEndpoint = m(eps.GetAllUnAuthEndpoint)
	}
	for _, m := range mdw["PostAuthInfo"] {
		eps.PostAuthInfoEndpoint = m(eps.PostAuthInfoEndpoint)
	}
	for _, m := range mdw["PostCertifyInfo"] {
		eps.PostCertifyInfoEndpoint = m(eps.PostCertifyInfoEndpoint)
	}
	return eps
}