// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

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

// EditBonusBalanceHandlerFunc turns a function with the right signature into a edit bonus balance handler
type EditBonusBalanceHandlerFunc func(EditBonusBalanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EditBonusBalanceHandlerFunc) Handle(params EditBonusBalanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EditBonusBalanceHandler interface for that can handle valid edit bonus balance params
type EditBonusBalanceHandler interface {
	Handle(EditBonusBalanceParams, interface{}) middleware.Responder
}

// NewEditBonusBalance creates a new http.Handler for the edit bonus balance operation
func NewEditBonusBalance(ctx *middleware.Context, handler EditBonusBalanceHandler) *EditBonusBalance {
	return &EditBonusBalance{Context: ctx, Handler: handler}
}

/*
	EditBonusBalance swagger:route PUT /balance/edit BonusBalance editBonusBalance

EditBonusBalance edit bonus balance API
*/
type EditBonusBalance struct {
	Context *middleware.Context
	Handler EditBonusBalanceHandler
}

func (o *EditBonusBalance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEditBonusBalanceParams()
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

// EditBonusBalanceBody edit bonus balance body
//
// swagger:model EditBonusBalanceBody
type EditBonusBalanceBody struct {

	// data
	Data *models.BalanceAdd `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this edit bonus balance body
func (o *EditBonusBalanceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditBonusBalanceBody) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this edit bonus balance body based on the context it is used
func (o *EditBonusBalanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditBonusBalanceBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *EditBonusBalanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditBonusBalanceBody) UnmarshalBinary(b []byte) error {
	var res EditBonusBalanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
