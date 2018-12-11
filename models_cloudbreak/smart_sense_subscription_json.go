// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmartSenseSubscriptionJSON smart sense subscription Json
// swagger:model SmartSenseSubscriptionJson

type SmartSenseSubscriptionJSON struct {

	// account id of the resource owner that is provided by OAuth provider
	// Read Only: true
	Account string `json:"account,omitempty"`

	// Flag of aut generated SmartSense subscription.
	AutoGenerated *bool `json:"autoGenerated,omitempty"`

	// id of the resource
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// id of the resource owner that is provided by OAuth provider
	// Read Only: true
	Owner string `json:"owner,omitempty"`

	// resource is visible in account
	// Read Only: true
	PublicInAccount *bool `json:"publicInAccount,omitempty"`

	// Identifier of SmartSense subscription.
	// Required: true
	// Pattern: ^([a-zA-Z]{1}-[0-9]{8}-[a-zA-Z]{1}-[0-9]{8}$)
	SubscriptionID *string `json:"subscriptionId"`
}

/* polymorph SmartSenseSubscriptionJson account false */

/* polymorph SmartSenseSubscriptionJson autoGenerated false */

/* polymorph SmartSenseSubscriptionJson id false */

/* polymorph SmartSenseSubscriptionJson owner false */

/* polymorph SmartSenseSubscriptionJson publicInAccount false */

/* polymorph SmartSenseSubscriptionJson subscriptionId false */

// Validate validates this smart sense subscription Json
func (m *SmartSenseSubscriptionJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartSenseSubscriptionJSON) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	if err := validate.Pattern("subscriptionId", "body", string(*m.SubscriptionID), `^([a-zA-Z]{1}-[0-9]{8}-[a-zA-Z]{1}-[0-9]{8}$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmartSenseSubscriptionJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartSenseSubscriptionJSON) UnmarshalBinary(b []byte) error {
	var res SmartSenseSubscriptionJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}