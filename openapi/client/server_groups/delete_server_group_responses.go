// Code generated by go-swagger; DO NOT EDIT.

package server_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washBonus/openapi/models"
)

// DeleteServerGroupReader is a Reader for the DeleteServerGroup structure.
type DeleteServerGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServerGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServerGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServerGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServerGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServerGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /server-groups/{groupId}] deleteServerGroup", response, response.Code())
	}
}

// NewDeleteServerGroupNoContent creates a DeleteServerGroupNoContent with default headers values
func NewDeleteServerGroupNoContent() *DeleteServerGroupNoContent {
	return &DeleteServerGroupNoContent{}
}

/*
DeleteServerGroupNoContent describes a response with status code 204, with default header values.

OK
*/
type DeleteServerGroupNoContent struct {
}

// IsSuccess returns true when this delete server group no content response has a 2xx status code
func (o *DeleteServerGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete server group no content response has a 3xx status code
func (o *DeleteServerGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete server group no content response has a 4xx status code
func (o *DeleteServerGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete server group no content response has a 5xx status code
func (o *DeleteServerGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete server group no content response a status code equal to that given
func (o *DeleteServerGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete server group no content response
func (o *DeleteServerGroupNoContent) Code() int {
	return 204
}

func (o *DeleteServerGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupNoContent ", 204)
}

func (o *DeleteServerGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupNoContent ", 204)
}

func (o *DeleteServerGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerGroupBadRequest creates a DeleteServerGroupBadRequest with default headers values
func NewDeleteServerGroupBadRequest() *DeleteServerGroupBadRequest {
	return &DeleteServerGroupBadRequest{}
}

/*
DeleteServerGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteServerGroupBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete server group bad request response has a 2xx status code
func (o *DeleteServerGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete server group bad request response has a 3xx status code
func (o *DeleteServerGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete server group bad request response has a 4xx status code
func (o *DeleteServerGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete server group bad request response has a 5xx status code
func (o *DeleteServerGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete server group bad request response a status code equal to that given
func (o *DeleteServerGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete server group bad request response
func (o *DeleteServerGroupBadRequest) Code() int {
	return 400
}

func (o *DeleteServerGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServerGroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServerGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServerGroupForbidden creates a DeleteServerGroupForbidden with default headers values
func NewDeleteServerGroupForbidden() *DeleteServerGroupForbidden {
	return &DeleteServerGroupForbidden{}
}

/*
DeleteServerGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteServerGroupForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete server group forbidden response has a 2xx status code
func (o *DeleteServerGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete server group forbidden response has a 3xx status code
func (o *DeleteServerGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete server group forbidden response has a 4xx status code
func (o *DeleteServerGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete server group forbidden response has a 5xx status code
func (o *DeleteServerGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete server group forbidden response a status code equal to that given
func (o *DeleteServerGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete server group forbidden response
func (o *DeleteServerGroupForbidden) Code() int {
	return 403
}

func (o *DeleteServerGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServerGroupForbidden) String() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteServerGroupForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServerGroupNotFound creates a DeleteServerGroupNotFound with default headers values
func NewDeleteServerGroupNotFound() *DeleteServerGroupNotFound {
	return &DeleteServerGroupNotFound{}
}

/*
DeleteServerGroupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteServerGroupNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete server group not found response has a 2xx status code
func (o *DeleteServerGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete server group not found response has a 3xx status code
func (o *DeleteServerGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete server group not found response has a 4xx status code
func (o *DeleteServerGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete server group not found response has a 5xx status code
func (o *DeleteServerGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete server group not found response a status code equal to that given
func (o *DeleteServerGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete server group not found response
func (o *DeleteServerGroupNotFound) Code() int {
	return 404
}

func (o *DeleteServerGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServerGroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServerGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServerGroupInternalServerError creates a DeleteServerGroupInternalServerError with default headers values
func NewDeleteServerGroupInternalServerError() *DeleteServerGroupInternalServerError {
	return &DeleteServerGroupInternalServerError{}
}

/*
DeleteServerGroupInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type DeleteServerGroupInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete server group internal server error response has a 2xx status code
func (o *DeleteServerGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete server group internal server error response has a 3xx status code
func (o *DeleteServerGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete server group internal server error response has a 4xx status code
func (o *DeleteServerGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete server group internal server error response has a 5xx status code
func (o *DeleteServerGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete server group internal server error response a status code equal to that given
func (o *DeleteServerGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete server group internal server error response
func (o *DeleteServerGroupInternalServerError) Code() int {
	return 500
}

func (o *DeleteServerGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServerGroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /server-groups/{groupId}][%d] deleteServerGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServerGroupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
