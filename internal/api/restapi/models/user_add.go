// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserAdd user model for add and edit methods
//
// swagger:model userAdd
type UserAdd struct {

	// active
	Active bool `json:"active,omitempty"`

	// firebase Id
	FirebaseID string `json:"firebaseId,omitempty"`
}

// Validate validates this user add
func (m *UserAdd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user add based on context it is used
func (m *UserAdd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAdd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAdd) UnmarshalBinary(b []byte) error {
	var res UserAdd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
