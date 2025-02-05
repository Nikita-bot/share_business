// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerGroup server group
//
// swagger:model ServerGroup
type ServerGroup struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organizationId,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *ServerGroup) UnmarshalJSON(data []byte) error {
	var props struct {

		// description
		Description string `json:"description,omitempty"`

		// id
		// Format: uuid
		ID strfmt.UUID `json:"id,omitempty"`

		// is default
		IsDefault bool `json:"isDefault,omitempty"`

		// name
		Name string `json:"name,omitempty"`

		// organization Id
		// Format: uuid
		OrganizationID strfmt.UUID `json:"organizationId,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Description = props.Description
	m.ID = props.ID
	m.IsDefault = props.IsDefault
	m.Name = props.Name
	m.OrganizationID = props.OrganizationID
	return nil
}

// Validate validates this server group
func (m *ServerGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroup) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerGroup) validateOrganizationID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organizationId", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server group based on context it is used
func (m *ServerGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerGroup) UnmarshalBinary(b []byte) error {
	var res ServerGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
