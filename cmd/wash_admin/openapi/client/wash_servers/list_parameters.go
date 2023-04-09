// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

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

	"wash_admin/openapi/models"
)

// NewListParams creates a new ListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListParams() *ListParams {
	return &ListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListParamsWithTimeout creates a new ListParams object
// with the ability to set a timeout on a request.
func NewListParamsWithTimeout(timeout time.Duration) *ListParams {
	return &ListParams{
		timeout: timeout,
	}
}

// NewListParamsWithContext creates a new ListParams object
// with the ability to set a context for a request.
func NewListParamsWithContext(ctx context.Context) *ListParams {
	return &ListParams{
		Context: ctx,
	}
}

// NewListParamsWithHTTPClient creates a new ListParams object
// with the ability to set a custom HTTPClient for a request.
func NewListParamsWithHTTPClient(client *http.Client) *ListParams {
	return &ListParams{
		HTTPClient: client,
	}
}

/*
ListParams contains all the parameters to send to the API endpoint

	for the list operation.

	Typically these are written to a http.Request.
*/
type ListParams struct {

	// Body.
	Body *models.Pagination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParams) WithDefaults() *ListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list params
func (o *ListParams) WithTimeout(timeout time.Duration) *ListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list params
func (o *ListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list params
func (o *ListParams) WithContext(ctx context.Context) *ListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list params
func (o *ListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list params
func (o *ListParams) WithHTTPClient(client *http.Client) *ListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list params
func (o *ListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list params
func (o *ListParams) WithBody(body *models.Pagination) *ListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list params
func (o *ListParams) SetBody(body *models.Pagination) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
