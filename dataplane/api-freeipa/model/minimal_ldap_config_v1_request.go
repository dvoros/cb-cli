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

// MinimalLdapConfigV1Request minimal ldap config v1 request
// swagger:model MinimalLdapConfigV1Request
type MinimalLdapConfigV1Request struct {

	// bind distinguished name for connection test and group search (e.g. cn=admin,dc=example,dc=org)
	// Required: true
	BindDn *string `json:"bindDn"`

	// password for the provided bind DN
	// Required: true
	BindPassword *string `json:"bindPassword"`

	// public host or IP address of LDAP server
	// Required: true
	Host *string `json:"host"`

	// port of LDAP server (typically: 389 or 636 for LDAPS)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// determines the protocol (LDAP or LDAP over SSL)
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this minimal ldap config v1 request
func (m *MinimalLdapConfigV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MinimalLdapConfigV1Request) validateBindDn(formats strfmt.Registry) error {

	if err := validate.Required("bindDn", "body", m.BindDn); err != nil {
		return err
	}

	return nil
}

func (m *MinimalLdapConfigV1Request) validateBindPassword(formats strfmt.Registry) error {

	if err := validate.Required("bindPassword", "body", m.BindPassword); err != nil {
		return err
	}

	return nil
}

func (m *MinimalLdapConfigV1Request) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *MinimalLdapConfigV1Request) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MinimalLdapConfigV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinimalLdapConfigV1Request) UnmarshalBinary(b []byte) error {
	var res MinimalLdapConfigV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}