// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetOrganizationsParams creates a new GetOrganizationsParams object
// with the default values initialized.
func NewGetOrganizationsParams() GetOrganizationsParams {

	var (
		// initialize parameters with default values

		isManagedByMeDefault = bool(false)
		limitDefault         = int64(100)
		offsetDefault        = int64(0)
	)

	return GetOrganizationsParams{
		IsManagedByMe: &isManagedByMeDefault,

		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetOrganizationsParams contains all the bound params for the get organizations operation
// typically these are obtained from a http.Request
//
// swagger:parameters getOrganizations
type GetOrganizationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Collection Format: csv
	*/
	Ids []strfmt.UUID
	/*
	  In: query
	  Default: false
	*/
	IsManagedByMe *bool
	/*Maximum number of records to return
	  Minimum: 0
	  In: query
	  Default: 100
	*/
	Limit *int64
	/*Number of records to skip for pagination
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetOrganizationsParams() beforehand.
func (o *GetOrganizationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIds, qhkIds, _ := qs.GetOK("ids")
	if err := o.bindIds(qIds, qhkIds, route.Formats); err != nil {
		res = append(res, err)
	}

	qIsManagedByMe, qhkIsManagedByMe, _ := qs.GetOK("isManagedByMe")
	if err := o.bindIsManagedByMe(qIsManagedByMe, qhkIsManagedByMe, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIds binds and validates array parameter Ids from query.
//
// Arrays are parsed according to CollectionFormat: "csv" (defaults to "csv" when empty).
func (o *GetOrganizationsParams) bindIds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvIds string
	if len(rawData) > 0 {
		qvIds = rawData[len(rawData)-1]
	}

	// CollectionFormat: csv
	idsIC := swag.SplitByFormat(qvIds, "csv")
	if len(idsIC) == 0 {
		return nil
	}

	var idsIR []strfmt.UUID
	for i, idsIV := range idsIC {
		// items.Format: "uuid"
		value, err := formats.Parse("uuid", idsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "ids", i), "query", "strfmt.UUID", value)
		}
		idsI := *(value.(*strfmt.UUID))

		if err := validate.FormatOf(fmt.Sprintf("%s.%v", "ids", i), "query", "uuid", idsI.String(), formats); err != nil {
			return err
		}
		idsIR = append(idsIR, idsI)
	}

	o.Ids = idsIR

	return nil
}

// bindIsManagedByMe binds and validates parameter IsManagedByMe from query.
func (o *GetOrganizationsParams) bindIsManagedByMe(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetOrganizationsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("isManagedByMe", "query", "bool", raw)
	}
	o.IsManagedByMe = &value

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetOrganizationsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetOrganizationsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetOrganizationsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", *o.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetOrganizationsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetOrganizationsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetOrganizationsParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", *o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}
