// Code generated by go-swagger; DO NOT EDIT.

package bonus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bonus API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bonus API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Cancel(params *CancelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelNoContent, error)

	Confirm(params *ConfirmParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConfirmNoContent, error)

	Use(params *UseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UseNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Cancel cancel API
*/
func (a *Client) Cancel(params *CancelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancel",
		Method:             "POST",
		PathPattern:        "/cancel-use",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Confirm confirm API
*/
func (a *Client) Confirm(params *ConfirmParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ConfirmNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfirmParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "confirm",
		Method:             "POST",
		PathPattern:        "/confirm-use",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConfirmReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfirmNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for confirm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Use use API
*/
func (a *Client) Use(params *UseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "use",
		Method:             "POST",
		PathPattern:        "/use-bonuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for use: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
