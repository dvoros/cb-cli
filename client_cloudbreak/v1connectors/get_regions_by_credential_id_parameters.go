// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

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

// NewGetRegionsByCredentialIDParams creates a new GetRegionsByCredentialIDParams object
// with the default values initialized.
func NewGetRegionsByCredentialIDParams() *GetRegionsByCredentialIDParams {
	var ()
	return &GetRegionsByCredentialIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsByCredentialIDParamsWithTimeout creates a new GetRegionsByCredentialIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsByCredentialIDParamsWithTimeout(timeout time.Duration) *GetRegionsByCredentialIDParams {
	var ()
	return &GetRegionsByCredentialIDParams{

		timeout: timeout,
	}
}

// NewGetRegionsByCredentialIDParamsWithContext creates a new GetRegionsByCredentialIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsByCredentialIDParamsWithContext(ctx context.Context) *GetRegionsByCredentialIDParams {
	var ()
	return &GetRegionsByCredentialIDParams{

		Context: ctx,
	}
}

// NewGetRegionsByCredentialIDParamsWithHTTPClient creates a new GetRegionsByCredentialIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsByCredentialIDParamsWithHTTPClient(client *http.Client) *GetRegionsByCredentialIDParams {
	var ()
	return &GetRegionsByCredentialIDParams{
		HTTPClient: client,
	}
}

/*GetRegionsByCredentialIDParams contains all the parameters to send to the API endpoint
for the get regions by credential Id operation typically these are written to a http.Request
*/
type GetRegionsByCredentialIDParams struct {

	/*Body*/
	Body *models_cloudbreak.PlatformResourceRequestJSON

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) WithTimeout(timeout time.Duration) *GetRegionsByCredentialIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) WithContext(ctx context.Context) *GetRegionsByCredentialIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) WithHTTPClient(client *http.Client) *GetRegionsByCredentialIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) WithBody(body *models_cloudbreak.PlatformResourceRequestJSON) *GetRegionsByCredentialIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get regions by credential Id params
func (o *GetRegionsByCredentialIDParams) SetBody(body *models_cloudbreak.PlatformResourceRequestJSON) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsByCredentialIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.PlatformResourceRequestJSON)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
