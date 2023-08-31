// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WashServerUpdate wash server update
//
// swagger:model WashServerUpdate
type WashServerUpdate struct {

	// description
	Description *string `json:"description,omitempty"`

	// name
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *WashServerUpdate) UnmarshalJSON(data []byte) error {
	var props struct {

		// description
		Description *string `json:"description,omitempty"`

		// name
		Name *string `json:"name,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Description = props.Description
	m.Name = props.Name
	return nil
}

// Validate validates this wash server update
func (m *WashServerUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wash server update based on context it is used
func (m *WashServerUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WashServerUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WashServerUpdate) UnmarshalBinary(b []byte) error {
	var res WashServerUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
