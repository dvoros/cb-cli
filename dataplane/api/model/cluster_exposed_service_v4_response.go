// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterExposedServiceV4Response cluster exposed service v4 response
// swagger:model ClusterExposedServiceV4Response
type ClusterExposedServiceV4Response struct {

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// knox service
	KnoxService string `json:"knoxService,omitempty"`

	// mode
	// Enum: [SSO_PROVIDER NONE]
	Mode string `json:"mode,omitempty"`

	// open
	Open bool `json:"open,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// service Url
	ServiceURL string `json:"serviceUrl,omitempty"`
}

// Validate validates this cluster exposed service v4 response
func (m *ClusterExposedServiceV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterExposedServiceV4ResponseTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSO_PROVIDER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterExposedServiceV4ResponseTypeModePropEnum = append(clusterExposedServiceV4ResponseTypeModePropEnum, v)
	}
}

const (

	// ClusterExposedServiceV4ResponseModeSSOPROVIDER captures enum value "SSO_PROVIDER"
	ClusterExposedServiceV4ResponseModeSSOPROVIDER string = "SSO_PROVIDER"

	// ClusterExposedServiceV4ResponseModeNONE captures enum value "NONE"
	ClusterExposedServiceV4ResponseModeNONE string = "NONE"
)

// prop value enum
func (m *ClusterExposedServiceV4Response) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterExposedServiceV4ResponseTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterExposedServiceV4Response) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterExposedServiceV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterExposedServiceV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterExposedServiceV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
