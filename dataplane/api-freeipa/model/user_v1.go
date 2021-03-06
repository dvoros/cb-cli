// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserV1 user v1
// swagger:model UserV1
type UserV1 struct {

	// first name of the user
	// Required: true
	FirstName *string `json:"firstName"`

	// groups to sync
	// Unique: true
	Groups []string `json:"groups"`

	// last name of the user
	// Required: true
	LastName *string `json:"lastName"`

	// name of the user
	// Required: true
	Name *string `json:"name"`

	// the user's password
	Password string `json:"password,omitempty"`
}

// Validate validates this user v1
func (m *UserV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
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

func (m *UserV1) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *UserV1) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	if err := validate.UniqueItems("groups", "body", m.Groups); err != nil {
		return err
	}

	return nil
}

func (m *UserV1) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *UserV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserV1) UnmarshalBinary(b []byte) error {
	var res UserV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
