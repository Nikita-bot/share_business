// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washBonus/openapi/models"
)

// GetServerGroupsOKCode is the HTTP code returned for type GetServerGroupsOK
const GetServerGroupsOKCode int = 200

/*
GetServerGroupsOK OK

swagger:response getServerGroupsOK
*/
type GetServerGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ServerGroup `json:"body,omitempty"`
}

// NewGetServerGroupsOK creates GetServerGroupsOK with default headers values
func NewGetServerGroupsOK() *GetServerGroupsOK {

	return &GetServerGroupsOK{}
}

// WithPayload adds the payload to the get server groups o k response
func (o *GetServerGroupsOK) WithPayload(payload []*models.ServerGroup) *GetServerGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server groups o k response
func (o *GetServerGroupsOK) SetPayload(payload []*models.ServerGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ServerGroup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *GetServerGroupsOK) GetServerGroupsResponder() {}

// GetServerGroupsBadRequestCode is the HTTP code returned for type GetServerGroupsBadRequest
const GetServerGroupsBadRequestCode int = 400

/*
GetServerGroupsBadRequest Bad request

swagger:response getServerGroupsBadRequest
*/
type GetServerGroupsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerGroupsBadRequest creates GetServerGroupsBadRequest with default headers values
func NewGetServerGroupsBadRequest() *GetServerGroupsBadRequest {

	return &GetServerGroupsBadRequest{}
}

// WithPayload adds the payload to the get server groups bad request response
func (o *GetServerGroupsBadRequest) WithPayload(payload *models.Error) *GetServerGroupsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server groups bad request response
func (o *GetServerGroupsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerGroupsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetServerGroupsBadRequest) GetServerGroupsResponder() {}

// GetServerGroupsForbiddenCode is the HTTP code returned for type GetServerGroupsForbidden
const GetServerGroupsForbiddenCode int = 403

/*
GetServerGroupsForbidden Forbidden

swagger:response getServerGroupsForbidden
*/
type GetServerGroupsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerGroupsForbidden creates GetServerGroupsForbidden with default headers values
func NewGetServerGroupsForbidden() *GetServerGroupsForbidden {

	return &GetServerGroupsForbidden{}
}

// WithPayload adds the payload to the get server groups forbidden response
func (o *GetServerGroupsForbidden) WithPayload(payload *models.Error) *GetServerGroupsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server groups forbidden response
func (o *GetServerGroupsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerGroupsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetServerGroupsForbidden) GetServerGroupsResponder() {}

// GetServerGroupsInternalServerErrorCode is the HTTP code returned for type GetServerGroupsInternalServerError
const GetServerGroupsInternalServerErrorCode int = 500

/*
GetServerGroupsInternalServerError Internal error

swagger:response getServerGroupsInternalServerError
*/
type GetServerGroupsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServerGroupsInternalServerError creates GetServerGroupsInternalServerError with default headers values
func NewGetServerGroupsInternalServerError() *GetServerGroupsInternalServerError {

	return &GetServerGroupsInternalServerError{}
}

// WithPayload adds the payload to the get server groups internal server error response
func (o *GetServerGroupsInternalServerError) WithPayload(payload *models.Error) *GetServerGroupsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server groups internal server error response
func (o *GetServerGroupsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerGroupsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *GetServerGroupsInternalServerError) GetServerGroupsResponder() {}

type GetServerGroupsNotImplementedResponder struct {
	middleware.Responder
}

func (*GetServerGroupsNotImplementedResponder) GetServerGroupsResponder() {}

func GetServerGroupsNotImplemented() GetServerGroupsResponder {
	return &GetServerGroupsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.GetServerGroups has not yet been implemented",
		),
	}
}

type GetServerGroupsResponder interface {
	middleware.Responder
	GetServerGroupsResponder()
}
