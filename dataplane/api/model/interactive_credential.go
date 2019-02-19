// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InteractiveCredential Contains values from an Azure interactive login attempt
// swagger:model InteractiveCredential
type InteractiveCredential struct {

	// The user code what has to be used for the sign-in process on the Azure portal
	UserCode string `json:"userCode,omitempty"`

	// The url provided by Azure where the user have to use the given user code to sign in
	VerificationURL string `json:"verificationUrl,omitempty"`
}

// Validate validates this interactive credential
func (m *InteractiveCredential) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InteractiveCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InteractiveCredential) UnmarshalBinary(b []byte) error {
	var res InteractiveCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}