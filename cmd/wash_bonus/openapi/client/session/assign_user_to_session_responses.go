// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash_bonus/openapi/models"
)

// AssignUserToSessionReader is a Reader for the AssignUserToSession structure.
type AssignUserToSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignUserToSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAssignUserToSessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAssignUserToSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssignUserToSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignUserToSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignUserToSessionNoContent creates a AssignUserToSessionNoContent with default headers values
func NewAssignUserToSessionNoContent() *AssignUserToSessionNoContent {
	return &AssignUserToSessionNoContent{}
}

/*
AssignUserToSessionNoContent describes a response with status code 204, with default header values.

OK
*/
type AssignUserToSessionNoContent struct {
}

// IsSuccess returns true when this assign user to session no content response has a 2xx status code
func (o *AssignUserToSessionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign user to session no content response has a 3xx status code
func (o *AssignUserToSessionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user to session no content response has a 4xx status code
func (o *AssignUserToSessionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign user to session no content response has a 5xx status code
func (o *AssignUserToSessionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this assign user to session no content response a status code equal to that given
func (o *AssignUserToSessionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the assign user to session no content response
func (o *AssignUserToSessionNoContent) Code() int {
	return 204
}

func (o *AssignUserToSessionNoContent) Error() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionNoContent ", 204)
}

func (o *AssignUserToSessionNoContent) String() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionNoContent ", 204)
}

func (o *AssignUserToSessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignUserToSessionForbidden creates a AssignUserToSessionForbidden with default headers values
func NewAssignUserToSessionForbidden() *AssignUserToSessionForbidden {
	return &AssignUserToSessionForbidden{}
}

/*
AssignUserToSessionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AssignUserToSessionForbidden struct {
}

// IsSuccess returns true when this assign user to session forbidden response has a 2xx status code
func (o *AssignUserToSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign user to session forbidden response has a 3xx status code
func (o *AssignUserToSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user to session forbidden response has a 4xx status code
func (o *AssignUserToSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign user to session forbidden response has a 5xx status code
func (o *AssignUserToSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign user to session forbidden response a status code equal to that given
func (o *AssignUserToSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the assign user to session forbidden response
func (o *AssignUserToSessionForbidden) Code() int {
	return 403
}

func (o *AssignUserToSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionForbidden ", 403)
}

func (o *AssignUserToSessionForbidden) String() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionForbidden ", 403)
}

func (o *AssignUserToSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignUserToSessionNotFound creates a AssignUserToSessionNotFound with default headers values
func NewAssignUserToSessionNotFound() *AssignUserToSessionNotFound {
	return &AssignUserToSessionNotFound{}
}

/*
AssignUserToSessionNotFound describes a response with status code 404, with default header values.

Session not found
*/
type AssignUserToSessionNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this assign user to session not found response has a 2xx status code
func (o *AssignUserToSessionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign user to session not found response has a 3xx status code
func (o *AssignUserToSessionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user to session not found response has a 4xx status code
func (o *AssignUserToSessionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign user to session not found response has a 5xx status code
func (o *AssignUserToSessionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this assign user to session not found response a status code equal to that given
func (o *AssignUserToSessionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the assign user to session not found response
func (o *AssignUserToSessionNotFound) Code() int {
	return 404
}

func (o *AssignUserToSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionNotFound  %+v", 404, o.Payload)
}

func (o *AssignUserToSessionNotFound) String() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionNotFound  %+v", 404, o.Payload)
}

func (o *AssignUserToSessionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignUserToSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignUserToSessionInternalServerError creates a AssignUserToSessionInternalServerError with default headers values
func NewAssignUserToSessionInternalServerError() *AssignUserToSessionInternalServerError {
	return &AssignUserToSessionInternalServerError{}
}

/*
AssignUserToSessionInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type AssignUserToSessionInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this assign user to session internal server error response has a 2xx status code
func (o *AssignUserToSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign user to session internal server error response has a 3xx status code
func (o *AssignUserToSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user to session internal server error response has a 4xx status code
func (o *AssignUserToSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign user to session internal server error response has a 5xx status code
func (o *AssignUserToSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assign user to session internal server error response a status code equal to that given
func (o *AssignUserToSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the assign user to session internal server error response
func (o *AssignUserToSessionInternalServerError) Code() int {
	return 500
}

func (o *AssignUserToSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignUserToSessionInternalServerError) String() string {
	return fmt.Sprintf("[POST /session/{sessionId}/assign-user][%d] assignUserToSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignUserToSessionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignUserToSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
