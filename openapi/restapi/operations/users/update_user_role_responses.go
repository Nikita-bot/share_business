// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washBonus/openapi/models"
)

// UpdateUserRoleNoContentCode is the HTTP code returned for type UpdateUserRoleNoContent
const UpdateUserRoleNoContentCode int = 204

/*
UpdateUserRoleNoContent OK

swagger:response updateUserRoleNoContent
*/
type UpdateUserRoleNoContent struct {
}

// NewUpdateUserRoleNoContent creates UpdateUserRoleNoContent with default headers values
func NewUpdateUserRoleNoContent() *UpdateUserRoleNoContent {

	return &UpdateUserRoleNoContent{}
}

// WriteResponse to the client
func (o *UpdateUserRoleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateUserRoleNoContent) UpdateUserRoleResponder() {}

// UpdateUserRoleBadRequestCode is the HTTP code returned for type UpdateUserRoleBadRequest
const UpdateUserRoleBadRequestCode int = 400

/*
UpdateUserRoleBadRequest Bad request

swagger:response updateUserRoleBadRequest
*/
type UpdateUserRoleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserRoleBadRequest creates UpdateUserRoleBadRequest with default headers values
func NewUpdateUserRoleBadRequest() *UpdateUserRoleBadRequest {

	return &UpdateUserRoleBadRequest{}
}

// WithPayload adds the payload to the update user role bad request response
func (o *UpdateUserRoleBadRequest) WithPayload(payload *models.Error) *UpdateUserRoleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user role bad request response
func (o *UpdateUserRoleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserRoleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserRoleBadRequest) UpdateUserRoleResponder() {}

// UpdateUserRoleForbiddenCode is the HTTP code returned for type UpdateUserRoleForbidden
const UpdateUserRoleForbiddenCode int = 403

/*
UpdateUserRoleForbidden Forbidden

swagger:response updateUserRoleForbidden
*/
type UpdateUserRoleForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserRoleForbidden creates UpdateUserRoleForbidden with default headers values
func NewUpdateUserRoleForbidden() *UpdateUserRoleForbidden {

	return &UpdateUserRoleForbidden{}
}

// WithPayload adds the payload to the update user role forbidden response
func (o *UpdateUserRoleForbidden) WithPayload(payload *models.Error) *UpdateUserRoleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user role forbidden response
func (o *UpdateUserRoleForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserRoleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserRoleForbidden) UpdateUserRoleResponder() {}

// UpdateUserRoleNotFoundCode is the HTTP code returned for type UpdateUserRoleNotFound
const UpdateUserRoleNotFoundCode int = 404

/*
UpdateUserRoleNotFound Not Found

swagger:response updateUserRoleNotFound
*/
type UpdateUserRoleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserRoleNotFound creates UpdateUserRoleNotFound with default headers values
func NewUpdateUserRoleNotFound() *UpdateUserRoleNotFound {

	return &UpdateUserRoleNotFound{}
}

// WithPayload adds the payload to the update user role not found response
func (o *UpdateUserRoleNotFound) WithPayload(payload *models.Error) *UpdateUserRoleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user role not found response
func (o *UpdateUserRoleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserRoleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserRoleNotFound) UpdateUserRoleResponder() {}

// UpdateUserRoleInternalServerErrorCode is the HTTP code returned for type UpdateUserRoleInternalServerError
const UpdateUserRoleInternalServerErrorCode int = 500

/*
UpdateUserRoleInternalServerError Internal error

swagger:response updateUserRoleInternalServerError
*/
type UpdateUserRoleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserRoleInternalServerError creates UpdateUserRoleInternalServerError with default headers values
func NewUpdateUserRoleInternalServerError() *UpdateUserRoleInternalServerError {

	return &UpdateUserRoleInternalServerError{}
}

// WithPayload adds the payload to the update user role internal server error response
func (o *UpdateUserRoleInternalServerError) WithPayload(payload *models.Error) *UpdateUserRoleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user role internal server error response
func (o *UpdateUserRoleInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserRoleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUserRoleInternalServerError) UpdateUserRoleResponder() {}

type UpdateUserRoleNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateUserRoleNotImplementedResponder) UpdateUserRoleResponder() {}

func UpdateUserRoleNotImplemented() UpdateUserRoleResponder {
	return &UpdateUserRoleNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateUserRole has not yet been implemented",
		),
	}
}

type UpdateUserRoleResponder interface {
	middleware.Responder
	UpdateUserRoleResponder()
}
