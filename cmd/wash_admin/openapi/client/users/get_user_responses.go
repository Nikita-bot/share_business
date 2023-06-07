// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash_admin/openapi/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*
GetUserOK describes a response with status code 200, with default header values.

OK
*/
type GetUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get user o k response has a 2xx status code
func (o *GetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user o k response has a 3xx status code
func (o *GetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user o k response has a 4xx status code
func (o *GetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user o k response has a 5xx status code
func (o *GetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user o k response a status code equal to that given
func (o *GetUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user o k response
func (o *GetUserOK) Code() int {
	return 200
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*
GetUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserForbidden struct {
}

// IsSuccess returns true when this get user forbidden response has a 2xx status code
func (o *GetUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user forbidden response has a 3xx status code
func (o *GetUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user forbidden response has a 4xx status code
func (o *GetUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user forbidden response has a 5xx status code
func (o *GetUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user forbidden response a status code equal to that given
func (o *GetUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user forbidden response
func (o *GetUserForbidden) Code() int {
	return 403
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserForbidden ", 403)
}

func (o *GetUserForbidden) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserForbidden ", 403)
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*
GetUserNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetUserNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user not found response has a 2xx status code
func (o *GetUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user not found response has a 3xx status code
func (o *GetUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user not found response has a 4xx status code
func (o *GetUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user not found response has a 5xx status code
func (o *GetUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user not found response a status code equal to that given
func (o *GetUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user not found response
func (o *GetUserNotFound) Code() int {
	return 404
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInternalServerError creates a GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {
	return &GetUserInternalServerError{}
}

/*
GetUserInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type GetUserInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get user internal server error response has a 2xx status code
func (o *GetUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user internal server error response has a 3xx status code
func (o *GetUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user internal server error response has a 4xx status code
func (o *GetUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user internal server error response has a 5xx status code
func (o *GetUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user internal server error response a status code equal to that given
func (o *GetUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user internal server error response
func (o *GetUserInternalServerError) Code() int {
	return 500
}

func (o *GetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
