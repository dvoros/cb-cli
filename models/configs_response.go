package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ConfigsResponse configs response

swagger:model ConfigsResponse
*/
type ConfigsResponse struct {

	/* response object

	Required: true
	Unique: true
	*/
	Inputs []*BlueprintInputJSON `json:"inputs"`
}

// Validate validates this configs response
func (m *ConfigsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigsResponse) validateInputs(formats strfmt.Registry) error {

	if err := validate.Required("inputs", "body", m.Inputs); err != nil {
		return err
	}

	if err := validate.UniqueItems("inputs", "body", m.Inputs); err != nil {
		return err
	}

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if err := m.Inputs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
