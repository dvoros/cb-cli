// Code generated by go-swagger; DO NOT EDIT.

package v1freeipausersync

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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewSynchronizeUsersV1Params creates a new SynchronizeUsersV1Params object
// with the default values initialized.
func NewSynchronizeUsersV1Params() *SynchronizeUsersV1Params {
	var ()
	return &SynchronizeUsersV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSynchronizeUsersV1ParamsWithTimeout creates a new SynchronizeUsersV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSynchronizeUsersV1ParamsWithTimeout(timeout time.Duration) *SynchronizeUsersV1Params {
	var ()
	return &SynchronizeUsersV1Params{

		timeout: timeout,
	}
}

// NewSynchronizeUsersV1ParamsWithContext creates a new SynchronizeUsersV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewSynchronizeUsersV1ParamsWithContext(ctx context.Context) *SynchronizeUsersV1Params {
	var ()
	return &SynchronizeUsersV1Params{

		Context: ctx,
	}
}

// NewSynchronizeUsersV1ParamsWithHTTPClient creates a new SynchronizeUsersV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSynchronizeUsersV1ParamsWithHTTPClient(client *http.Client) *SynchronizeUsersV1Params {
	var ()
	return &SynchronizeUsersV1Params{
		HTTPClient: client,
	}
}

/*SynchronizeUsersV1Params contains all the parameters to send to the API endpoint
for the synchronize users v1 operation typically these are written to a http.Request
*/
type SynchronizeUsersV1Params struct {

	/*Body*/
	Body *model.SynchronizeUsersV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) WithTimeout(timeout time.Duration) *SynchronizeUsersV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) WithContext(ctx context.Context) *SynchronizeUsersV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) WithHTTPClient(client *http.Client) *SynchronizeUsersV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) WithBody(body *model.SynchronizeUsersV1Request) *SynchronizeUsersV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the synchronize users v1 params
func (o *SynchronizeUsersV1Params) SetBody(body *model.SynchronizeUsersV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SynchronizeUsersV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
