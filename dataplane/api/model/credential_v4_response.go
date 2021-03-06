// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CredentialV4Response credential v4 response
// swagger:model CredentialV4Response
type CredentialV4Response struct {
	CredentialV4Base

	// provider specific attributes of the credential
	Attributes *SecretResponse `json:"attributes,omitempty"`

	// custom parameters for Azure credential
	Azure *AzureCredentialV4ResponseParameters `json:"azure,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// workspace
	Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CredentialV4Response) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CredentialV4Base
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CredentialV4Base = aO0

	// AO1
	var dataAO1 struct {
		Attributes *SecretResponse `json:"attributes,omitempty"`

		Azure *AzureCredentialV4ResponseParameters `json:"azure,omitempty"`

		ID int64 `json:"id,omitempty"`

		Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Attributes = dataAO1.Attributes

	m.Azure = dataAO1.Azure

	m.ID = dataAO1.ID

	m.Workspace = dataAO1.Workspace

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CredentialV4Response) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CredentialV4Base)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Attributes *SecretResponse `json:"attributes,omitempty"`

		Azure *AzureCredentialV4ResponseParameters `json:"azure,omitempty"`

		ID int64 `json:"id,omitempty"`

		Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
	}

	dataAO1.Attributes = m.Attributes

	dataAO1.Azure = m.Azure

	dataAO1.ID = m.ID

	dataAO1.Workspace = m.Workspace

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this credential v4 response
func (m *CredentialV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CredentialV4Base
	if err := m.CredentialV4Base.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialV4Response) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialV4Response) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialV4Response) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialV4Response) UnmarshalBinary(b []byte) error {
	var res CredentialV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
