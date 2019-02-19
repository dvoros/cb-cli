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

// CumulusYarnCredentialV4Parameters cumulus yarn credential v4 parameters
// swagger:model CumulusYarnCredentialV4Parameters
type CumulusYarnCredentialV4Parameters struct {

	// ambari password
	// Required: true
	AmbariPassword *string `json:"ambariPassword"`

	// ambari Url
	// Required: true
	AmbariURL *string `json:"ambariUrl"`

	// ambari user
	// Required: true
	AmbariUser *string `json:"ambariUser"`
}

// Validate validates this cumulus yarn credential v4 parameters
func (m *CumulusYarnCredentialV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmbariURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmbariUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CumulusYarnCredentialV4Parameters) validateAmbariPassword(formats strfmt.Registry) error {

	if err := validate.Required("ambariPassword", "body", m.AmbariPassword); err != nil {
		return err
	}

	return nil
}

func (m *CumulusYarnCredentialV4Parameters) validateAmbariURL(formats strfmt.Registry) error {

	if err := validate.Required("ambariUrl", "body", m.AmbariURL); err != nil {
		return err
	}

	return nil
}

func (m *CumulusYarnCredentialV4Parameters) validateAmbariUser(formats strfmt.Registry) error {

	if err := validate.Required("ambariUser", "body", m.AmbariUser); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CumulusYarnCredentialV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CumulusYarnCredentialV4Parameters) UnmarshalBinary(b []byte) error {
	var res CumulusYarnCredentialV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}