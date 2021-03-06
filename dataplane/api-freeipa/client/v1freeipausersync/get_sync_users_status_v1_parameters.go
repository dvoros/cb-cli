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
)

// NewGetSyncUsersStatusV1Params creates a new GetSyncUsersStatusV1Params object
// with the default values initialized.
func NewGetSyncUsersStatusV1Params() *GetSyncUsersStatusV1Params {

	return &GetSyncUsersStatusV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncUsersStatusV1ParamsWithTimeout creates a new GetSyncUsersStatusV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSyncUsersStatusV1ParamsWithTimeout(timeout time.Duration) *GetSyncUsersStatusV1Params {

	return &GetSyncUsersStatusV1Params{

		timeout: timeout,
	}
}

// NewGetSyncUsersStatusV1ParamsWithContext creates a new GetSyncUsersStatusV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetSyncUsersStatusV1ParamsWithContext(ctx context.Context) *GetSyncUsersStatusV1Params {

	return &GetSyncUsersStatusV1Params{

		Context: ctx,
	}
}

// NewGetSyncUsersStatusV1ParamsWithHTTPClient creates a new GetSyncUsersStatusV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSyncUsersStatusV1ParamsWithHTTPClient(client *http.Client) *GetSyncUsersStatusV1Params {

	return &GetSyncUsersStatusV1Params{
		HTTPClient: client,
	}
}

/*GetSyncUsersStatusV1Params contains all the parameters to send to the API endpoint
for the get sync users status v1 operation typically these are written to a http.Request
*/
type GetSyncUsersStatusV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) WithTimeout(timeout time.Duration) *GetSyncUsersStatusV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) WithContext(ctx context.Context) *GetSyncUsersStatusV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) WithHTTPClient(client *http.Client) *GetSyncUsersStatusV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync users status v1 params
func (o *GetSyncUsersStatusV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncUsersStatusV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
