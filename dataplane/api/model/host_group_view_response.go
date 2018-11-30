// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostGroupViewResponse host group view response
// swagger:model HostGroupViewResponse
type HostGroupViewResponse struct {

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// metadata of hosts
	// Unique: true
	Metadata []*HostMetadataViewResponse `json:"metadata"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this host group view response
func (m *HostGroupViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupViewResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.UniqueItems("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostGroupViewResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostGroupViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostGroupViewResponse) UnmarshalBinary(b []byte) error {
	var res HostGroupViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}