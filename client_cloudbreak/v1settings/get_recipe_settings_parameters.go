// Code generated by go-swagger; DO NOT EDIT.

package v1settings

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

// NewGetRecipeSettingsParams creates a new GetRecipeSettingsParams object
// with the default values initialized.
func NewGetRecipeSettingsParams() *GetRecipeSettingsParams {

	return &GetRecipeSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecipeSettingsParamsWithTimeout creates a new GetRecipeSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecipeSettingsParamsWithTimeout(timeout time.Duration) *GetRecipeSettingsParams {

	return &GetRecipeSettingsParams{

		timeout: timeout,
	}
}

// NewGetRecipeSettingsParamsWithContext creates a new GetRecipeSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecipeSettingsParamsWithContext(ctx context.Context) *GetRecipeSettingsParams {

	return &GetRecipeSettingsParams{

		Context: ctx,
	}
}

// NewGetRecipeSettingsParamsWithHTTPClient creates a new GetRecipeSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecipeSettingsParamsWithHTTPClient(client *http.Client) *GetRecipeSettingsParams {

	return &GetRecipeSettingsParams{
		HTTPClient: client,
	}
}

/*GetRecipeSettingsParams contains all the parameters to send to the API endpoint
for the get recipe settings operation typically these are written to a http.Request
*/
type GetRecipeSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recipe settings params
func (o *GetRecipeSettingsParams) WithTimeout(timeout time.Duration) *GetRecipeSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recipe settings params
func (o *GetRecipeSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recipe settings params
func (o *GetRecipeSettingsParams) WithContext(ctx context.Context) *GetRecipeSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recipe settings params
func (o *GetRecipeSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recipe settings params
func (o *GetRecipeSettingsParams) WithHTTPClient(client *http.Client) *GetRecipeSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recipe settings params
func (o *GetRecipeSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecipeSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
