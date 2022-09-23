// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"wash-bonus/internal/api/restapi/models"
)

// NewAddBonusBalanceParams creates a new AddBonusBalanceParams object
//
// There are no default values defined in the spec.
func NewAddBonusBalanceParams() AddBonusBalanceParams {

	return AddBonusBalanceParams{}
}

// AddBonusBalanceParams contains all the bound params for the add bonus balance operation
// typically these are obtained from a http.Request
//
// swagger:parameters addBonusBalance
type AddBonusBalanceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.BalanceAdd
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddBonusBalanceParams() beforehand.
func (o *AddBonusBalanceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.BalanceAdd
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
