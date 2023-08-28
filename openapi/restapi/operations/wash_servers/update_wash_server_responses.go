// Code generated by go-swagger; DO NOT EDIT.

package wash_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washBonus/openapi/models"
)

// UpdateWashServerOKCode is the HTTP code returned for type UpdateWashServerOK
const UpdateWashServerOKCode int = 200

/*
UpdateWashServerOK Success update

swagger:response updateWashServerOK
*/
type UpdateWashServerOK struct {

	/*
	  In: Body
	*/
	Payload *models.WashServer `json:"body,omitempty"`
}

// NewUpdateWashServerOK creates UpdateWashServerOK with default headers values
func NewUpdateWashServerOK() *UpdateWashServerOK {

	return &UpdateWashServerOK{}
}

// WithPayload adds the payload to the update wash server o k response
func (o *UpdateWashServerOK) WithPayload(payload *models.WashServer) *UpdateWashServerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash server o k response
func (o *UpdateWashServerOK) SetPayload(payload *models.WashServer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashServerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashServerOK) UpdateWashServerResponder() {}

// UpdateWashServerBadRequestCode is the HTTP code returned for type UpdateWashServerBadRequest
const UpdateWashServerBadRequestCode int = 400

/*
UpdateWashServerBadRequest Bad request

swagger:response updateWashServerBadRequest
*/
type UpdateWashServerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateWashServerBadRequest creates UpdateWashServerBadRequest with default headers values
func NewUpdateWashServerBadRequest() *UpdateWashServerBadRequest {

	return &UpdateWashServerBadRequest{}
}

// WithPayload adds the payload to the update wash server bad request response
func (o *UpdateWashServerBadRequest) WithPayload(payload *models.Error) *UpdateWashServerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash server bad request response
func (o *UpdateWashServerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashServerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashServerBadRequest) UpdateWashServerResponder() {}

// UpdateWashServerForbiddenCode is the HTTP code returned for type UpdateWashServerForbidden
const UpdateWashServerForbiddenCode int = 403

/*
UpdateWashServerForbidden Forbidden

swagger:response updateWashServerForbidden
*/
type UpdateWashServerForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateWashServerForbidden creates UpdateWashServerForbidden with default headers values
func NewUpdateWashServerForbidden() *UpdateWashServerForbidden {

	return &UpdateWashServerForbidden{}
}

// WithPayload adds the payload to the update wash server forbidden response
func (o *UpdateWashServerForbidden) WithPayload(payload *models.Error) *UpdateWashServerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash server forbidden response
func (o *UpdateWashServerForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashServerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashServerForbidden) UpdateWashServerResponder() {}

// UpdateWashServerNotFoundCode is the HTTP code returned for type UpdateWashServerNotFound
const UpdateWashServerNotFoundCode int = 404

/*
UpdateWashServerNotFound Not Found

swagger:response updateWashServerNotFound
*/
type UpdateWashServerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateWashServerNotFound creates UpdateWashServerNotFound with default headers values
func NewUpdateWashServerNotFound() *UpdateWashServerNotFound {

	return &UpdateWashServerNotFound{}
}

// WithPayload adds the payload to the update wash server not found response
func (o *UpdateWashServerNotFound) WithPayload(payload *models.Error) *UpdateWashServerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash server not found response
func (o *UpdateWashServerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashServerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashServerNotFound) UpdateWashServerResponder() {}

// UpdateWashServerInternalServerErrorCode is the HTTP code returned for type UpdateWashServerInternalServerError
const UpdateWashServerInternalServerErrorCode int = 500

/*
UpdateWashServerInternalServerError Internal error

swagger:response updateWashServerInternalServerError
*/
type UpdateWashServerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateWashServerInternalServerError creates UpdateWashServerInternalServerError with default headers values
func NewUpdateWashServerInternalServerError() *UpdateWashServerInternalServerError {

	return &UpdateWashServerInternalServerError{}
}

// WithPayload adds the payload to the update wash server internal server error response
func (o *UpdateWashServerInternalServerError) WithPayload(payload *models.Error) *UpdateWashServerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update wash server internal server error response
func (o *UpdateWashServerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateWashServerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateWashServerInternalServerError) UpdateWashServerResponder() {}

type UpdateWashServerNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateWashServerNotImplementedResponder) UpdateWashServerResponder() {}

func UpdateWashServerNotImplemented() UpdateWashServerResponder {
	return &UpdateWashServerNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateWashServer has not yet been implemented",
		),
	}
}

type UpdateWashServerResponder interface {
	middleware.Responder
	UpdateWashServerResponder()
}
