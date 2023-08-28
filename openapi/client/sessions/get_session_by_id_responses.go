// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"washBonus/openapi/models"
)

// GetSessionByIDReader is a Reader for the GetSessionByID structure.
type GetSessionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSessionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSessionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSessionByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSessionByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSessionByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSessionByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sessions/{sessionId}] getSessionById", response, response.Code())
	}
}

// NewGetSessionByIDOK creates a GetSessionByIDOK with default headers values
func NewGetSessionByIDOK() *GetSessionByIDOK {
	return &GetSessionByIDOK{}
}

/*
GetSessionByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetSessionByIDOK struct {
	Payload *models.Session
}

// IsSuccess returns true when this get session by Id o k response has a 2xx status code
func (o *GetSessionByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get session by Id o k response has a 3xx status code
func (o *GetSessionByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id o k response has a 4xx status code
func (o *GetSessionByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session by Id o k response has a 5xx status code
func (o *GetSessionByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get session by Id o k response a status code equal to that given
func (o *GetSessionByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get session by Id o k response
func (o *GetSessionByIDOK) Code() int {
	return 200
}

func (o *GetSessionByIDOK) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdOK  %+v", 200, o.Payload)
}

func (o *GetSessionByIDOK) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdOK  %+v", 200, o.Payload)
}

func (o *GetSessionByIDOK) GetPayload() *models.Session {
	return o.Payload
}

func (o *GetSessionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionByIDBadRequest creates a GetSessionByIDBadRequest with default headers values
func NewGetSessionByIDBadRequest() *GetSessionByIDBadRequest {
	return &GetSessionByIDBadRequest{}
}

/*
GetSessionByIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSessionByIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session by Id bad request response has a 2xx status code
func (o *GetSessionByIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session by Id bad request response has a 3xx status code
func (o *GetSessionByIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id bad request response has a 4xx status code
func (o *GetSessionByIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session by Id bad request response has a 5xx status code
func (o *GetSessionByIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get session by Id bad request response a status code equal to that given
func (o *GetSessionByIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get session by Id bad request response
func (o *GetSessionByIDBadRequest) Code() int {
	return 400
}

func (o *GetSessionByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSessionByIDBadRequest) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetSessionByIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionByIDForbidden creates a GetSessionByIDForbidden with default headers values
func NewGetSessionByIDForbidden() *GetSessionByIDForbidden {
	return &GetSessionByIDForbidden{}
}

/*
GetSessionByIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSessionByIDForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session by Id forbidden response has a 2xx status code
func (o *GetSessionByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session by Id forbidden response has a 3xx status code
func (o *GetSessionByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id forbidden response has a 4xx status code
func (o *GetSessionByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session by Id forbidden response has a 5xx status code
func (o *GetSessionByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get session by Id forbidden response a status code equal to that given
func (o *GetSessionByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get session by Id forbidden response
func (o *GetSessionByIDForbidden) Code() int {
	return 403
}

func (o *GetSessionByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetSessionByIDForbidden) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetSessionByIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionByIDNotFound creates a GetSessionByIDNotFound with default headers values
func NewGetSessionByIDNotFound() *GetSessionByIDNotFound {
	return &GetSessionByIDNotFound{}
}

/*
GetSessionByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSessionByIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session by Id not found response has a 2xx status code
func (o *GetSessionByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session by Id not found response has a 3xx status code
func (o *GetSessionByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id not found response has a 4xx status code
func (o *GetSessionByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get session by Id not found response has a 5xx status code
func (o *GetSessionByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get session by Id not found response a status code equal to that given
func (o *GetSessionByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get session by Id not found response
func (o *GetSessionByIDNotFound) Code() int {
	return 404
}

func (o *GetSessionByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSessionByIDNotFound) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSessionByIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSessionByIDInternalServerError creates a GetSessionByIDInternalServerError with default headers values
func NewGetSessionByIDInternalServerError() *GetSessionByIDInternalServerError {
	return &GetSessionByIDInternalServerError{}
}

/*
GetSessionByIDInternalServerError describes a response with status code 500, with default header values.

Internal error
*/
type GetSessionByIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get session by Id internal server error response has a 2xx status code
func (o *GetSessionByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get session by Id internal server error response has a 3xx status code
func (o *GetSessionByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get session by Id internal server error response has a 4xx status code
func (o *GetSessionByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get session by Id internal server error response has a 5xx status code
func (o *GetSessionByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get session by Id internal server error response a status code equal to that given
func (o *GetSessionByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get session by Id internal server error response
func (o *GetSessionByIDInternalServerError) Code() int {
	return 500
}

func (o *GetSessionByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /sessions/{sessionId}][%d] getSessionByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSessionByIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSessionByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
