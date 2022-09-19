// Code generated by go-swagger; DO NOT EDIT.

package wash_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wash-bonus/internal/api/restapi/models"
)

// AddWashSessionReader is a Reader for the AddWashSession structure.
type AddWashSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddWashSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddWashSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddWashSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddWashSessionCreated creates a AddWashSessionCreated with default headers values
func NewAddWashSessionCreated() *AddWashSessionCreated {
	return &AddWashSessionCreated{}
}

/*
AddWashSessionCreated describes a response with status code 201, with default header values.

Created
*/
type AddWashSessionCreated struct {
	Payload *models.WashSession
}

// IsSuccess returns true when this add wash session created response has a 2xx status code
func (o *AddWashSessionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add wash session created response has a 3xx status code
func (o *AddWashSessionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add wash session created response has a 4xx status code
func (o *AddWashSessionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add wash session created response has a 5xx status code
func (o *AddWashSessionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add wash session created response a status code equal to that given
func (o *AddWashSessionCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddWashSessionCreated) Error() string {
	return fmt.Sprintf("[POST /washSession/add][%d] addWashSessionCreated  %+v", 201, o.Payload)
}

func (o *AddWashSessionCreated) String() string {
	return fmt.Sprintf("[POST /washSession/add][%d] addWashSessionCreated  %+v", 201, o.Payload)
}

func (o *AddWashSessionCreated) GetPayload() *models.WashSession {
	return o.Payload
}

func (o *AddWashSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WashSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddWashSessionDefault creates a AddWashSessionDefault with default headers values
func NewAddWashSessionDefault(code int) *AddWashSessionDefault {
	return &AddWashSessionDefault{
		_statusCode: code,
	}
}

/*
AddWashSessionDefault describes a response with status code -1, with default header values.

error
*/
type AddWashSessionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add wash session default response
func (o *AddWashSessionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this add wash session default response has a 2xx status code
func (o *AddWashSessionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add wash session default response has a 3xx status code
func (o *AddWashSessionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add wash session default response has a 4xx status code
func (o *AddWashSessionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add wash session default response has a 5xx status code
func (o *AddWashSessionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add wash session default response a status code equal to that given
func (o *AddWashSessionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AddWashSessionDefault) Error() string {
	return fmt.Sprintf("[POST /washSession/add][%d] addWashSession default  %+v", o._statusCode, o.Payload)
}

func (o *AddWashSessionDefault) String() string {
	return fmt.Sprintf("[POST /washSession/add][%d] addWashSession default  %+v", o._statusCode, o.Payload)
}

func (o *AddWashSessionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddWashSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
