// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteWashSessionHandlerFunc turns a function with the right signature into a delete wash session handler
type DeleteWashSessionHandlerFunc func(DeleteWashSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWashSessionHandlerFunc) Handle(params DeleteWashSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteWashSessionHandler interface for that can handle valid delete wash session params
type DeleteWashSessionHandler interface {
	Handle(DeleteWashSessionParams, interface{}) middleware.Responder
}

// NewDeleteWashSession creates a new http.Handler for the delete wash session operation
func NewDeleteWashSession(ctx *middleware.Context, handler DeleteWashSessionHandler) *DeleteWashSession {
	return &DeleteWashSession{Context: ctx, Handler: handler}
}

/*
	DeleteWashSession swagger:route DELETE /washSession/delete WashSession deleteWashSession

DeleteWashSession delete wash session API
*/
type DeleteWashSession struct {
	Context *middleware.Context
	Handler DeleteWashSessionHandler
}

func (o *DeleteWashSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteWashSessionParams()
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

// DeleteWashSessionBody delete wash session body
//
// swagger:model DeleteWashSessionBody
type DeleteWashSessionBody struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this delete wash session body
func (o *DeleteWashSessionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete wash session body based on context it is used
func (o *DeleteWashSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWashSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWashSessionBody) UnmarshalBinary(b []byte) error {
	var res DeleteWashSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
