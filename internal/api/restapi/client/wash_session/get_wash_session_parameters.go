// Code generated by go-swagger; DO NOT EDIT.

package wash_session

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
)

// NewGetWashSessionParams creates a new GetWashSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWashSessionParams() *GetWashSessionParams {
	return &GetWashSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWashSessionParamsWithTimeout creates a new GetWashSessionParams object
// with the ability to set a timeout on a request.
func NewGetWashSessionParamsWithTimeout(timeout time.Duration) *GetWashSessionParams {
	return &GetWashSessionParams{
		timeout: timeout,
	}
}

// NewGetWashSessionParamsWithContext creates a new GetWashSessionParams object
// with the ability to set a context for a request.
func NewGetWashSessionParamsWithContext(ctx context.Context) *GetWashSessionParams {
	return &GetWashSessionParams{
		Context: ctx,
	}
}

// NewGetWashSessionParamsWithHTTPClient creates a new GetWashSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWashSessionParamsWithHTTPClient(client *http.Client) *GetWashSessionParams {
	return &GetWashSessionParams{
		HTTPClient: client,
	}
}

/*
GetWashSessionParams contains all the parameters to send to the API endpoint

	for the get wash session operation.

	Typically these are written to a http.Request.
*/
type GetWashSessionParams struct {

	// Body.
	Body GetWashSessionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wash session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashSessionParams) WithDefaults() *GetWashSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wash session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWashSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wash session params
func (o *GetWashSessionParams) WithTimeout(timeout time.Duration) *GetWashSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wash session params
func (o *GetWashSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wash session params
func (o *GetWashSessionParams) WithContext(ctx context.Context) *GetWashSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wash session params
func (o *GetWashSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wash session params
func (o *GetWashSessionParams) WithHTTPClient(client *http.Client) *GetWashSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wash session params
func (o *GetWashSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get wash session params
func (o *GetWashSessionParams) WithBody(body GetWashSessionBody) *GetWashSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get wash session params
func (o *GetWashSessionParams) SetBody(body GetWashSessionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetWashSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
