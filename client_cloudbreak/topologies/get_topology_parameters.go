package topologies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetTopologyParams creates a new GetTopologyParams object
// with the default values initialized.
func NewGetTopologyParams() *GetTopologyParams {
	var ()
	return &GetTopologyParams{}
}

/*GetTopologyParams contains all the parameters to send to the API endpoint
for the get topology operation typically these are written to a http.Request
*/
type GetTopologyParams struct {

	/*ID*/
	ID int64
}

// WithID adds the id to the get topology params
func (o *GetTopologyParams) WithID(id int64) *GetTopologyParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTopologyParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}