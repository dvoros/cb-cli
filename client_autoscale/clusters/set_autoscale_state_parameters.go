package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// NewSetAutoscaleStateParams creates a new SetAutoscaleStateParams object
// with the default values initialized.
func NewSetAutoscaleStateParams() *SetAutoscaleStateParams {
	var ()
	return &SetAutoscaleStateParams{}
}

/*SetAutoscaleStateParams contains all the parameters to send to the API endpoint
for the set autoscale state operation typically these are written to a http.Request
*/
type SetAutoscaleStateParams struct {

	/*Body*/
	Body *models_autoscale.ClusterAutoscaleState
	/*ClusterID*/
	ClusterID int64
}

// WithBody adds the body to the set autoscale state params
func (o *SetAutoscaleStateParams) WithBody(body *models_autoscale.ClusterAutoscaleState) *SetAutoscaleStateParams {
	o.Body = body
	return o
}

// WithClusterID adds the clusterId to the set autoscale state params
func (o *SetAutoscaleStateParams) WithClusterID(clusterId int64) *SetAutoscaleStateParams {
	o.ClusterID = clusterId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SetAutoscaleStateParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ClusterAutoscaleState)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}