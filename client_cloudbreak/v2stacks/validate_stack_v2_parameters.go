// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

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

// NewValidateStackV2Params creates a new ValidateStackV2Params object
// with the default values initialized.
func NewValidateStackV2Params() *ValidateStackV2Params {
	var ()
	return &ValidateStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateStackV2ParamsWithTimeout creates a new ValidateStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateStackV2ParamsWithTimeout(timeout time.Duration) *ValidateStackV2Params {
	var ()
	return &ValidateStackV2Params{

		timeout: timeout,
	}
}

// NewValidateStackV2ParamsWithContext creates a new ValidateStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewValidateStackV2ParamsWithContext(ctx context.Context) *ValidateStackV2Params {
	var ()
	return &ValidateStackV2Params{

		Context: ctx,
	}
}

// NewValidateStackV2ParamsWithHTTPClient creates a new ValidateStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateStackV2ParamsWithHTTPClient(client *http.Client) *ValidateStackV2Params {
	var ()
	return &ValidateStackV2Params{
		HTTPClient: client,
	}
}

/*ValidateStackV2Params contains all the parameters to send to the API endpoint
for the validate stack v2 operation typically these are written to a http.Request
*/
type ValidateStackV2Params struct {

	/*Body*/
	Body *models_cloudbreak.StackValidationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate stack v2 params
func (o *ValidateStackV2Params) WithTimeout(timeout time.Duration) *ValidateStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate stack v2 params
func (o *ValidateStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate stack v2 params
func (o *ValidateStackV2Params) WithContext(ctx context.Context) *ValidateStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate stack v2 params
func (o *ValidateStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate stack v2 params
func (o *ValidateStackV2Params) WithHTTPClient(client *http.Client) *ValidateStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate stack v2 params
func (o *ValidateStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate stack v2 params
func (o *ValidateStackV2Params) WithBody(body *models_cloudbreak.StackValidationRequest) *ValidateStackV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate stack v2 params
func (o *ValidateStackV2Params) SetBody(body *models_cloudbreak.StackValidationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.StackValidationRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
