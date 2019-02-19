// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AccessConfigV4Response access config v4 response
// swagger:model AccessConfigV4Response
type AccessConfigV4Response struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Validate validates this access config v4 response
func (m *AccessConfigV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessConfigV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessConfigV4Response) UnmarshalBinary(b []byte) error {
	var res AccessConfigV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}