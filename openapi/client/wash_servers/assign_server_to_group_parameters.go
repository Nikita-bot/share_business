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
)

// NewAssignServerToGroupParams creates a new AssignServerToGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignServerToGroupParams() *AssignServerToGroupParams {
	return &AssignServerToGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignServerToGroupParamsWithTimeout creates a new AssignServerToGroupParams object
// with the ability to set a timeout on a request.
func NewAssignServerToGroupParamsWithTimeout(timeout time.Duration) *AssignServerToGroupParams {
	return &AssignServerToGroupParams{
		timeout: timeout,
	}
}

// NewAssignServerToGroupParamsWithContext creates a new AssignServerToGroupParams object
// with the ability to set a context for a request.
func NewAssignServerToGroupParamsWithContext(ctx context.Context) *AssignServerToGroupParams {
	return &AssignServerToGroupParams{
		Context: ctx,
	}
}

// NewAssignServerToGroupParamsWithHTTPClient creates a new AssignServerToGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignServerToGroupParamsWithHTTPClient(client *http.Client) *AssignServerToGroupParams {
	return &AssignServerToGroupParams{
		HTTPClient: client,
	}
}

/*
AssignServerToGroupParams contains all the parameters to send to the API endpoint

	for the assign server to group operation.

	Typically these are written to a http.Request.
*/
type AssignServerToGroupParams struct {

	// GroupID.
	//
	// Format: uuid
	GroupID strfmt.UUID

	// ServerID.
	//
	// Format: uuid
	ServerID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign server to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignServerToGroupParams) WithDefaults() *AssignServerToGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign server to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignServerToGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign server to group params
func (o *AssignServerToGroupParams) WithTimeout(timeout time.Duration) *AssignServerToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign server to group params
func (o *AssignServerToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign server to group params
func (o *AssignServerToGroupParams) WithContext(ctx context.Context) *AssignServerToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign server to group params
func (o *AssignServerToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign server to group params
func (o *AssignServerToGroupParams) WithHTTPClient(client *http.Client) *AssignServerToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign server to group params
func (o *AssignServerToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the assign server to group params
func (o *AssignServerToGroupParams) WithGroupID(groupID strfmt.UUID) *AssignServerToGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the assign server to group params
func (o *AssignServerToGroupParams) SetGroupID(groupID strfmt.UUID) {
	o.GroupID = groupID
}

// WithServerID adds the serverID to the assign server to group params
func (o *AssignServerToGroupParams) WithServerID(serverID strfmt.UUID) *AssignServerToGroupParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the assign server to group params
func (o *AssignServerToGroupParams) SetServerID(serverID strfmt.UUID) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *AssignServerToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID.String()); err != nil {
		return err
	}

	// path param serverId
	if err := r.SetPathParam("serverId", o.ServerID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
