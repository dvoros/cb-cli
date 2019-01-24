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

// DatalakeResourcesResponse datalake resources response
// swagger:model DatalakeResourcesResponse
type DatalakeResourcesResponse struct {

	// Ambari url
	AmbariURL string `json:"ambariUrl,omitempty"`

	// Kerberos config name for the cluster
	KerberosName string `json:"kerberosName,omitempty"`

	// LDAP config name for the cluster
	LdapName string `json:"ldapName,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	RdsNames []string `json:"rdsNames"`

	// Descriptors of the datalake services
	ServiceDescriptorMap map[string]ServiceDescriptorResponse `json:"serviceDescriptorMap,omitempty"`
}

// Validate validates this datalake resources response
func (m *DatalakeResourcesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRdsNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDescriptorMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatalakeResourcesResponse) validateRdsNames(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsNames", "body", m.RdsNames); err != nil {
		return err
	}

	return nil
}

func (m *DatalakeResourcesResponse) validateServiceDescriptorMap(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceDescriptorMap) { // not required
		return nil
	}

	for k := range m.ServiceDescriptorMap {

		if err := validate.Required("serviceDescriptorMap"+"."+k, "body", m.ServiceDescriptorMap[k]); err != nil {
			return err
		}
		if val, ok := m.ServiceDescriptorMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatalakeResourcesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatalakeResourcesResponse) UnmarshalBinary(b []byte) error {
	var res DatalakeResourcesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}