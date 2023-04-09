// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new session API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for session API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssignUserToSession(params *AssignUserToSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssignUserToSessionNoContent, error)

	GetSession(params *GetSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSessionOK, error)

	PostSession(params *PostSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostSessionNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AssignUserToSession assign user to session API
*/
func (a *Client) AssignUserToSession(params *AssignUserToSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssignUserToSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignUserToSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assignUserToSession",
		Method:             "POST",
		PathPattern:        "/session/{sessionId}/assign-user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AssignUserToSessionReader{formats: a.formats},
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
	success, ok := result.(*AssignUserToSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assignUserToSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSession get session API
*/
func (a *Client) GetSession(params *GetSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSession",
		Method:             "GET",
		PathPattern:        "/session/{UID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSessionReader{formats: a.formats},
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
	success, ok := result.(*GetSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostSession post session API
*/
func (a *Client) PostSession(params *PostSessionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSessionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postSession",
		Method:             "POST",
		PathPattern:        "/session/{UID}/bonuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSessionReader{formats: a.formats},
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
	success, ok := result.(*PostSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
