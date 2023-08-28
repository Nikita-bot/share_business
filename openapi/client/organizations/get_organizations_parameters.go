// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"washBonus/openapi/models"
)

// NewGetOrganizationsParams creates a new GetOrganizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsParams() *GetOrganizationsParams {
	return &GetOrganizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsParamsWithTimeout creates a new GetOrganizationsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsParamsWithContext creates a new GetOrganizationsParams object
// with the ability to set a context for a request.
func NewGetOrganizationsParamsWithContext(ctx context.Context) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		Context: ctx,
	}
}

// NewGetOrganizationsParamsWithHTTPClient creates a new GetOrganizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsParamsWithHTTPClient(client *http.Client) *GetOrganizationsParams {
	return &GetOrganizationsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsParams contains all the parameters to send to the API endpoint

	for the get organizations operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsParams struct {

	// Body.
	Body *models.Pagination

	// Ids.
	Ids []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsParams) WithDefaults() *GetOrganizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations params
func (o *GetOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations params
func (o *GetOrganizationsParams) WithContext(ctx context.Context) *GetOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations params
func (o *GetOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) WithHTTPClient(client *http.Client) *GetOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations params
func (o *GetOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get organizations params
func (o *GetOrganizationsParams) WithBody(body *models.Pagination) *GetOrganizationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get organizations params
func (o *GetOrganizationsParams) SetBody(body *models.Pagination) {
	o.Body = body
}

// WithIds adds the ids to the get organizations params
func (o *GetOrganizationsParams) WithIds(ids []strfmt.UUID) *GetOrganizationsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get organizations params
func (o *GetOrganizationsParams) SetIds(ids []strfmt.UUID) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizations binds the parameter ids
func (o *GetOrganizationsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []strfmt.UUID

		idsIIV := idsIIR.String() // strfmt.UUID as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}
