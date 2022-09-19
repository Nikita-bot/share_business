// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// EditWashSessionHandlerFunc turns a function with the right signature into a edit wash session handler
type EditWashSessionHandlerFunc func(EditWashSessionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EditWashSessionHandlerFunc) Handle(params EditWashSessionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EditWashSessionHandler interface for that can handle valid edit wash session params
type EditWashSessionHandler interface {
	Handle(EditWashSessionParams, interface{}) middleware.Responder
}

// NewEditWashSession creates a new http.Handler for the edit wash session operation
func NewEditWashSession(ctx *middleware.Context, handler EditWashSessionHandler) *EditWashSession {
	return &EditWashSession{Context: ctx, Handler: handler}
}

/*
	EditWashSession swagger:route PUT /washSession/edit WashSession editWashSession

EditWashSession edit wash session API
*/
type EditWashSession struct {
	Context *middleware.Context
	Handler EditWashSessionHandler
}

func (o *EditWashSession) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEditWashSessionParams()
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

// EditWashSessionBody edit wash session body
//
// swagger:model EditWashSessionBody
type EditWashSessionBody struct {

	// data
	Data *models.WashSessionAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit wash session body
func (o *EditWashSessionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashSessionBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this edit wash session body based on the context it is used
func (o *EditWashSessionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditWashSessionBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EditWashSessionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditWashSessionBody) UnmarshalBinary(b []byte) error {
	var res EditWashSessionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
