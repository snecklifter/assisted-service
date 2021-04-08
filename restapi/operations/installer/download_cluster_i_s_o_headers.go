// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DownloadClusterISOHeadersHandlerFunc turns a function with the right signature into a download cluster i s o headers handler
type DownloadClusterISOHeadersHandlerFunc func(DownloadClusterISOHeadersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadClusterISOHeadersHandlerFunc) Handle(params DownloadClusterISOHeadersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DownloadClusterISOHeadersHandler interface for that can handle valid download cluster i s o headers params
type DownloadClusterISOHeadersHandler interface {
	Handle(DownloadClusterISOHeadersParams, interface{}) middleware.Responder
}

// NewDownloadClusterISOHeaders creates a new http.Handler for the download cluster i s o headers operation
func NewDownloadClusterISOHeaders(ctx *middleware.Context, handler DownloadClusterISOHeadersHandler) *DownloadClusterISOHeaders {
	return &DownloadClusterISOHeaders{Context: ctx, Handler: handler}
}

/*DownloadClusterISOHeaders swagger:route HEAD /clusters/{cluster_id}/downloads/image installer downloadClusterISOHeaders

Downloads the OpenShift per-cluster Discovery ISO Headers only.

*/
type DownloadClusterISOHeaders struct {
	Context *middleware.Context
	Handler DownloadClusterISOHeadersHandler
}

func (o *DownloadClusterISOHeaders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownloadClusterISOHeadersParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
