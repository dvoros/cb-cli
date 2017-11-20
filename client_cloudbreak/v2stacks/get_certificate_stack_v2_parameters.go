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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCertificateStackV2Params creates a new GetCertificateStackV2Params object
// with the default values initialized.
func NewGetCertificateStackV2Params() *GetCertificateStackV2Params {
	var ()
	return &GetCertificateStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateStackV2ParamsWithTimeout creates a new GetCertificateStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCertificateStackV2ParamsWithTimeout(timeout time.Duration) *GetCertificateStackV2Params {
	var ()
	return &GetCertificateStackV2Params{

		timeout: timeout,
	}
}

// NewGetCertificateStackV2ParamsWithContext creates a new GetCertificateStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCertificateStackV2ParamsWithContext(ctx context.Context) *GetCertificateStackV2Params {
	var ()
	return &GetCertificateStackV2Params{

		Context: ctx,
	}
}

// NewGetCertificateStackV2ParamsWithHTTPClient creates a new GetCertificateStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCertificateStackV2ParamsWithHTTPClient(client *http.Client) *GetCertificateStackV2Params {
	var ()
	return &GetCertificateStackV2Params{
		HTTPClient: client,
	}
}

/*GetCertificateStackV2Params contains all the parameters to send to the API endpoint
for the get certificate stack v2 operation typically these are written to a http.Request
*/
type GetCertificateStackV2Params struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) WithTimeout(timeout time.Duration) *GetCertificateStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) WithContext(ctx context.Context) *GetCertificateStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) WithHTTPClient(client *http.Client) *GetCertificateStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) WithID(id int64) *GetCertificateStackV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate stack v2 params
func (o *GetCertificateStackV2Params) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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