// Code generated by go-swagger; DO NOT EDIT.

package v1util

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewTestAmbariDatabaseUtilParams creates a new TestAmbariDatabaseUtilParams object
// with the default values initialized.
func NewTestAmbariDatabaseUtilParams() *TestAmbariDatabaseUtilParams {
	var ()
	return &TestAmbariDatabaseUtilParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestAmbariDatabaseUtilParamsWithTimeout creates a new TestAmbariDatabaseUtilParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestAmbariDatabaseUtilParamsWithTimeout(timeout time.Duration) *TestAmbariDatabaseUtilParams {
	var ()
	return &TestAmbariDatabaseUtilParams{

		timeout: timeout,
	}
}

// NewTestAmbariDatabaseUtilParamsWithContext creates a new TestAmbariDatabaseUtilParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestAmbariDatabaseUtilParamsWithContext(ctx context.Context) *TestAmbariDatabaseUtilParams {
	var ()
	return &TestAmbariDatabaseUtilParams{

		Context: ctx,
	}
}

// NewTestAmbariDatabaseUtilParamsWithHTTPClient creates a new TestAmbariDatabaseUtilParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestAmbariDatabaseUtilParamsWithHTTPClient(client *http.Client) *TestAmbariDatabaseUtilParams {
	var ()
	return &TestAmbariDatabaseUtilParams{
		HTTPClient: client,
	}
}

/*TestAmbariDatabaseUtilParams contains all the parameters to send to the API endpoint
for the test ambari database util operation typically these are written to a http.Request
*/
type TestAmbariDatabaseUtilParams struct {

	/*Body*/
	Body *models_cloudbreak.AmbariDatabaseDetails

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) WithTimeout(timeout time.Duration) *TestAmbariDatabaseUtilParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) WithContext(ctx context.Context) *TestAmbariDatabaseUtilParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) WithHTTPClient(client *http.Client) *TestAmbariDatabaseUtilParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) WithBody(body *models_cloudbreak.AmbariDatabaseDetails) *TestAmbariDatabaseUtilParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ambari database util params
func (o *TestAmbariDatabaseUtilParams) SetBody(body *models_cloudbreak.AmbariDatabaseDetails) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestAmbariDatabaseUtilParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.AmbariDatabaseDetails)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
