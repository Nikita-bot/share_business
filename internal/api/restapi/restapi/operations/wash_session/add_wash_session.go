// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddWashSessionHandlerFunc turns a function with the right signature into a add wash session handler
type AddWashSessionHandlerFunc func(AddWashSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddWashSessionHandlerFunc) Handle(params AddWashSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddWashSessionHandler interface for that can handle valid add wash session params
type AddWashSessionHandler interface {
	Handle(AddWashSessionParams, interface{}) middleware.Responder
}

// NewAddWashSession creates a new http.Handler for the add wash session operation
func NewAddWashSession(ctx *middleware.Context, handler AddWashSessionHandler) *AddWashSession {
	return &AddWashSession{Context: ctx, Handler: handler}
}

/*
	AddWashSession swagger:route POST /washSession/add WashSession addWashSession

AddWashSession add wash session API
*/
type AddWashSession struct {
	Context *middleware.Context
	Handler AddWashSessionHandler
}

func (o *AddWashSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddWashSessionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
