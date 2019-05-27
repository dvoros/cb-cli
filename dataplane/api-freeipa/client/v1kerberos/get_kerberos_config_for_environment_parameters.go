// Code generated by go-swagger; DO NOT EDIT.

package v1kerberos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKerberosConfigForEnvironmentParams creates a new GetKerberosConfigForEnvironmentParams object
// with the default values initialized.
func NewGetKerberosConfigForEnvironmentParams() *GetKerberosConfigForEnvironmentParams {
	var ()
	return &GetKerberosConfigForEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKerberosConfigForEnvironmentParamsWithTimeout creates a new GetKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKerberosConfigForEnvironmentParamsWithTimeout(timeout time.Duration) *GetKerberosConfigForEnvironmentParams {
	var ()
	return &GetKerberosConfigForEnvironmentParams{

		timeout: timeout,
	}
}

// NewGetKerberosConfigForEnvironmentParamsWithContext creates a new GetKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKerberosConfigForEnvironmentParamsWithContext(ctx context.Context) *GetKerberosConfigForEnvironmentParams {
	var ()
	return &GetKerberosConfigForEnvironmentParams{

		Context: ctx,
	}
}

// NewGetKerberosConfigForEnvironmentParamsWithHTTPClient creates a new GetKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKerberosConfigForEnvironmentParamsWithHTTPClient(client *http.Client) *GetKerberosConfigForEnvironmentParams {
	var ()
	return &GetKerberosConfigForEnvironmentParams{
		HTTPClient: client,
	}
}

/*GetKerberosConfigForEnvironmentParams contains all the parameters to send to the API endpoint
for the get kerberos config for environment operation typically these are written to a http.Request
*/
type GetKerberosConfigForEnvironmentParams struct {

	/*EnvironmentID*/
	EnvironmentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) WithTimeout(timeout time.Duration) *GetKerberosConfigForEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) WithContext(ctx context.Context) *GetKerberosConfigForEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) WithHTTPClient(client *http.Client) *GetKerberosConfigForEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) WithEnvironmentID(environmentID *string) *GetKerberosConfigForEnvironmentParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the get kerberos config for environment params
func (o *GetKerberosConfigForEnvironmentParams) SetEnvironmentID(environmentID *string) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKerberosConfigForEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentID != nil {

		// query param environmentId
		var qrEnvironmentID string
		if o.EnvironmentID != nil {
			qrEnvironmentID = *o.EnvironmentID
		}
		qEnvironmentID := qrEnvironmentID
		if qEnvironmentID != "" {
			if err := r.SetQueryParam("environmentId", qEnvironmentID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
