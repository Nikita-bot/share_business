// Code generated by go-swagger; DO NOT EDIT.

package server_groups

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
)

// NewGetServerGroupsParams creates a new GetServerGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerGroupsParams() *GetServerGroupsParams {
	return &GetServerGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerGroupsParamsWithTimeout creates a new GetServerGroupsParams object
// with the ability to set a timeout on a request.
func NewGetServerGroupsParamsWithTimeout(timeout time.Duration) *GetServerGroupsParams {
	return &GetServerGroupsParams{
		timeout: timeout,
	}
}

// NewGetServerGroupsParamsWithContext creates a new GetServerGroupsParams object
// with the ability to set a context for a request.
func NewGetServerGroupsParamsWithContext(ctx context.Context) *GetServerGroupsParams {
	return &GetServerGroupsParams{
		Context: ctx,
	}
}

// NewGetServerGroupsParamsWithHTTPClient creates a new GetServerGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerGroupsParamsWithHTTPClient(client *http.Client) *GetServerGroupsParams {
	return &GetServerGroupsParams{
		HTTPClient: client,
	}
}

/*
GetServerGroupsParams contains all the parameters to send to the API endpoint

	for the get server groups operation.

	Typically these are written to a http.Request.
*/
type GetServerGroupsParams struct {

	// IsManagedByMe.
	IsManagedByMe *bool

	/* Limit.

	   Maximum number of records to return

	   Format: int64
	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   Number of records to skip for pagination

	   Format: int64
	*/
	Offset *int64

	// OrganizationID.
	//
	// Format: uuid
	OrganizationID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupsParams) WithDefaults() *GetServerGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerGroupsParams) SetDefaults() {
	var (
		isManagedByMeDefault = bool(false)

		limitDefault = int64(100)

		offsetDefault = int64(0)
	)

	val := GetServerGroupsParams{
		IsManagedByMe: &isManagedByMeDefault,
		Limit:         &limitDefault,
		Offset:        &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get server groups params
func (o *GetServerGroupsParams) WithTimeout(timeout time.Duration) *GetServerGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server groups params
func (o *GetServerGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server groups params
func (o *GetServerGroupsParams) WithContext(ctx context.Context) *GetServerGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server groups params
func (o *GetServerGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server groups params
func (o *GetServerGroupsParams) WithHTTPClient(client *http.Client) *GetServerGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server groups params
func (o *GetServerGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsManagedByMe adds the isManagedByMe to the get server groups params
func (o *GetServerGroupsParams) WithIsManagedByMe(isManagedByMe *bool) *GetServerGroupsParams {
	o.SetIsManagedByMe(isManagedByMe)
	return o
}

// SetIsManagedByMe adds the isManagedByMe to the get server groups params
func (o *GetServerGroupsParams) SetIsManagedByMe(isManagedByMe *bool) {
	o.IsManagedByMe = isManagedByMe
}

// WithLimit adds the limit to the get server groups params
func (o *GetServerGroupsParams) WithLimit(limit *int64) *GetServerGroupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get server groups params
func (o *GetServerGroupsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get server groups params
func (o *GetServerGroupsParams) WithOffset(offset *int64) *GetServerGroupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get server groups params
func (o *GetServerGroupsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get server groups params
func (o *GetServerGroupsParams) WithOrganizationID(organizationID *strfmt.UUID) *GetServerGroupsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get server groups params
func (o *GetServerGroupsParams) SetOrganizationID(organizationID *strfmt.UUID) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsManagedByMe != nil {

		// query param isManagedByMe
		var qrIsManagedByMe bool

		if o.IsManagedByMe != nil {
			qrIsManagedByMe = *o.IsManagedByMe
		}
		qIsManagedByMe := swag.FormatBool(qrIsManagedByMe)
		if qIsManagedByMe != "" {

			if err := r.SetQueryParam("isManagedByMe", qIsManagedByMe); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID strfmt.UUID

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID.String()
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
