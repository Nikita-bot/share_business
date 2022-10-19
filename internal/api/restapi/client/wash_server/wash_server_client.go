// Code generated by go-swagger; DO NOT EDIT.

package wash_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wash server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wash server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddWashServer(params *AddWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddWashServerOK, error)

	DeleteWashServer(params *DeleteWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteWashServerNoContent, error)

	EditWashServer(params *EditWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditWashServerOK, error)

	GenerateKeyWashServer(params *GenerateKeyWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateKeyWashServerOK, error)

	GetWashServer(params *GetWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWashServerOK, error)

	ListWashServer(params *ListWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListWashServerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddWashServer add wash server API
*/
func (a *Client) AddWashServer(params *AddWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddWashServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addWashServer",
		Method:             "POST",
		PathPattern:        "/washServer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddWashServerReader{formats: a.formats},
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
	success, ok := result.(*AddWashServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteWashServer delete wash server API
*/
func (a *Client) DeleteWashServer(params *DeleteWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteWashServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteWashServer",
		Method:             "DELETE",
		PathPattern:        "/washServer/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWashServerReader{formats: a.formats},
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
	success, ok := result.(*DeleteWashServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  EditWashServer edit wash server API
*/
func (a *Client) EditWashServer(params *EditWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditWashServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editWashServer",
		Method:             "PUT",
		PathPattern:        "/washServer/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EditWashServerReader{formats: a.formats},
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
	success, ok := result.(*EditWashServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EditWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GenerateKeyWashServer generate key wash server API
*/
func (a *Client) GenerateKeyWashServer(params *GenerateKeyWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateKeyWashServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateKeyWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generateKeyWashServer",
		Method:             "PUT",
		PathPattern:        "/washServer/{id}/generate-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateKeyWashServerReader{formats: a.formats},
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
	success, ok := result.(*GenerateKeyWashServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GenerateKeyWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetWashServer get wash server API
*/
func (a *Client) GetWashServer(params *GetWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWashServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getWashServer",
		Method:             "GET",
		PathPattern:        "/washServer/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWashServerReader{formats: a.formats},
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
	success, ok := result.(*GetWashServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListWashServer list wash server API
*/
func (a *Client) ListWashServer(params *ListWashServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListWashServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListWashServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listWashServer",
		Method:             "POST",
		PathPattern:        "/washServers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListWashServerReader{formats: a.formats},
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
	success, ok := result.(*ListWashServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListWashServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
