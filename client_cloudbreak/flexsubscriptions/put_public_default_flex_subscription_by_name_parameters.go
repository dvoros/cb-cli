package flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewPutPublicDefaultFlexSubscriptionByNameParams creates a new PutPublicDefaultFlexSubscriptionByNameParams object
// with the default values initialized.
func NewPutPublicDefaultFlexSubscriptionByNameParams() *PutPublicDefaultFlexSubscriptionByNameParams {
	var ()
	return &PutPublicDefaultFlexSubscriptionByNameParams{}
}

/*PutPublicDefaultFlexSubscriptionByNameParams contains all the parameters to send to the API endpoint
for the put public default flex subscription by name operation typically these are written to a http.Request
*/
type PutPublicDefaultFlexSubscriptionByNameParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the put public default flex subscription by name params
func (o *PutPublicDefaultFlexSubscriptionByNameParams) WithName(name string) *PutPublicDefaultFlexSubscriptionByNameParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutPublicDefaultFlexSubscriptionByNameParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}