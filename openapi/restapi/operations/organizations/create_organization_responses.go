// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"washBonus/openapi/models"
)

// CreateOrganizationOKCode is the HTTP code returned for type CreateOrganizationOK
const CreateOrganizationOKCode int = 200

/*
CreateOrganizationOK Successfull created

swagger:response createOrganizationOK
*/
type CreateOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.Organization `json:"body,omitempty"`
}

// NewCreateOrganizationOK creates CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {

	return &CreateOrganizationOK{}
}

// WithPayload adds the payload to the create organization o k response
func (o *CreateOrganizationOK) WithPayload(payload *models.Organization) *CreateOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create organization o k response
func (o *CreateOrganizationOK) SetPayload(payload *models.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateOrganizationOK) CreateOrganizationResponder() {}

// CreateOrganizationBadRequestCode is the HTTP code returned for type CreateOrganizationBadRequest
const CreateOrganizationBadRequestCode int = 400

/*
CreateOrganizationBadRequest Bad request

swagger:response createOrganizationBadRequest
*/
type CreateOrganizationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrganizationBadRequest creates CreateOrganizationBadRequest with default headers values
func NewCreateOrganizationBadRequest() *CreateOrganizationBadRequest {

	return &CreateOrganizationBadRequest{}
}

// WithPayload adds the payload to the create organization bad request response
func (o *CreateOrganizationBadRequest) WithPayload(payload *models.Error) *CreateOrganizationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create organization bad request response
func (o *CreateOrganizationBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrganizationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateOrganizationBadRequest) CreateOrganizationResponder() {}

// CreateOrganizationForbiddenCode is the HTTP code returned for type CreateOrganizationForbidden
const CreateOrganizationForbiddenCode int = 403

/*
CreateOrganizationForbidden Forbidden

swagger:response createOrganizationForbidden
*/
type CreateOrganizationForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrganizationForbidden creates CreateOrganizationForbidden with default headers values
func NewCreateOrganizationForbidden() *CreateOrganizationForbidden {

	return &CreateOrganizationForbidden{}
}

// WithPayload adds the payload to the create organization forbidden response
func (o *CreateOrganizationForbidden) WithPayload(payload *models.Error) *CreateOrganizationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create organization forbidden response
func (o *CreateOrganizationForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrganizationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateOrganizationForbidden) CreateOrganizationResponder() {}

// CreateOrganizationInternalServerErrorCode is the HTTP code returned for type CreateOrganizationInternalServerError
const CreateOrganizationInternalServerErrorCode int = 500

/*
CreateOrganizationInternalServerError Internal error

swagger:response createOrganizationInternalServerError
*/
type CreateOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateOrganizationInternalServerError creates CreateOrganizationInternalServerError with default headers values
func NewCreateOrganizationInternalServerError() *CreateOrganizationInternalServerError {

	return &CreateOrganizationInternalServerError{}
}

// WithPayload adds the payload to the create organization internal server error response
func (o *CreateOrganizationInternalServerError) WithPayload(payload *models.Error) *CreateOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create organization internal server error response
func (o *CreateOrganizationInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CreateOrganizationInternalServerError) CreateOrganizationResponder() {}

type CreateOrganizationNotImplementedResponder struct {
	middleware.Responder
}

func (*CreateOrganizationNotImplementedResponder) CreateOrganizationResponder() {}

func CreateOrganizationNotImplemented() CreateOrganizationResponder {
	return &CreateOrganizationNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CreateOrganization has not yet been implemented",
		),
	}
}

type CreateOrganizationResponder interface {
	middleware.Responder
	CreateOrganizationResponder()
}
