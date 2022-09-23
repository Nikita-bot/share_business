// Code generated by go-swagger; DO NOT EDIT.

package bonus_balance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
)

// DeleteBonusBalanceReader is a Reader for the DeleteBonusBalance structure.
type DeleteBonusBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBonusBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBonusBalanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBonusBalanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBonusBalanceNoContent creates a DeleteBonusBalanceNoContent with default headers values
func NewDeleteBonusBalanceNoContent() *DeleteBonusBalanceNoContent {
	return &DeleteBonusBalanceNoContent{}
}

/*
DeleteBonusBalanceNoContent describes a response with status code 204, with default header values.

Deleted
*/
type DeleteBonusBalanceNoContent struct {
}

// IsSuccess returns true when this delete bonus balance no content response has a 2xx status code
func (o *DeleteBonusBalanceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete bonus balance no content response has a 3xx status code
func (o *DeleteBonusBalanceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete bonus balance no content response has a 4xx status code
func (o *DeleteBonusBalanceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete bonus balance no content response has a 5xx status code
func (o *DeleteBonusBalanceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete bonus balance no content response a status code equal to that given
func (o *DeleteBonusBalanceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteBonusBalanceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /balance/deleted][%d] deleteBonusBalanceNoContent ", 204)
}

func (o *DeleteBonusBalanceNoContent) String() string {
	return fmt.Sprintf("[DELETE /balance/deleted][%d] deleteBonusBalanceNoContent ", 204)
}

func (o *DeleteBonusBalanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBonusBalanceDefault creates a DeleteBonusBalanceDefault with default headers values
func NewDeleteBonusBalanceDefault(code int) *DeleteBonusBalanceDefault {
	return &DeleteBonusBalanceDefault{
		_statusCode: code,
	}
}

/*
DeleteBonusBalanceDefault describes a response with status code -1, with default header values.

error
*/
type DeleteBonusBalanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete bonus balance default response
func (o *DeleteBonusBalanceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete bonus balance default response has a 2xx status code
func (o *DeleteBonusBalanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete bonus balance default response has a 3xx status code
func (o *DeleteBonusBalanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete bonus balance default response has a 4xx status code
func (o *DeleteBonusBalanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete bonus balance default response has a 5xx status code
func (o *DeleteBonusBalanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete bonus balance default response a status code equal to that given
func (o *DeleteBonusBalanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteBonusBalanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /balance/deleted][%d] deleteBonusBalance default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBonusBalanceDefault) String() string {
	return fmt.Sprintf("[DELETE /balance/deleted][%d] deleteBonusBalance default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBonusBalanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBonusBalanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeleteBonusBalanceBody delete bonus balance body
swagger:model DeleteBonusBalanceBody
*/
type DeleteBonusBalanceBody struct {

	// id
	ID string `json:"id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this delete bonus balance body
func (o *DeleteBonusBalanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete bonus balance body based on context it is used
func (o *DeleteBonusBalanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBonusBalanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBonusBalanceBody) UnmarshalBinary(b []byte) error {
	var res DeleteBonusBalanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
