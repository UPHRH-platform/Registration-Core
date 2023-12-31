// Code generated by go-swagger; DO NOT EDIT.

package download_file_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"bulk_issuance/swagger_gen/models"
)

// GetV1IDReportHandlerFunc turns a function with the right signature into a get v1 ID report handler
type GetV1IDReportHandlerFunc func(GetV1IDReportParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1IDReportHandlerFunc) Handle(params GetV1IDReportParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetV1IDReportHandler interface for that can handle valid get v1 ID report params
type GetV1IDReportHandler interface {
	Handle(GetV1IDReportParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetV1IDReport creates a new http.Handler for the get v1 ID report operation
func NewGetV1IDReport(ctx *middleware.Context, handler GetV1IDReportHandler) *GetV1IDReport {
	return &GetV1IDReport{Context: ctx, Handler: handler}
}

/*
	GetV1IDReport swagger:route GET /v1/{id}/report downloadFileReport getV1IdReport

download the success and error report of file uploaded
*/
type GetV1IDReport struct {
	Context *middleware.Context
	Handler GetV1IDReportHandler
}

func (o *GetV1IDReport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetV1IDReportParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
