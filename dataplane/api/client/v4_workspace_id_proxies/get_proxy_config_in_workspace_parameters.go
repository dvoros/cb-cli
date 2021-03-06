// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

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

// NewGetProxyConfigInWorkspaceParams creates a new GetProxyConfigInWorkspaceParams object
// with the default values initialized.
func NewGetProxyConfigInWorkspaceParams() *GetProxyConfigInWorkspaceParams {
	var ()
	return &GetProxyConfigInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyConfigInWorkspaceParamsWithTimeout creates a new GetProxyConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProxyConfigInWorkspaceParamsWithTimeout(timeout time.Duration) *GetProxyConfigInWorkspaceParams {
	var ()
	return &GetProxyConfigInWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetProxyConfigInWorkspaceParamsWithContext creates a new GetProxyConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProxyConfigInWorkspaceParamsWithContext(ctx context.Context) *GetProxyConfigInWorkspaceParams {
	var ()
	return &GetProxyConfigInWorkspaceParams{

		Context: ctx,
	}
}

// NewGetProxyConfigInWorkspaceParamsWithHTTPClient creates a new GetProxyConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProxyConfigInWorkspaceParamsWithHTTPClient(client *http.Client) *GetProxyConfigInWorkspaceParams {
	var ()
	return &GetProxyConfigInWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetProxyConfigInWorkspaceParams contains all the parameters to send to the API endpoint
for the get proxy config in workspace operation typically these are written to a http.Request
*/
type GetProxyConfigInWorkspaceParams struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) WithTimeout(timeout time.Duration) *GetProxyConfigInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) WithContext(ctx context.Context) *GetProxyConfigInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) WithHTTPClient(client *http.Client) *GetProxyConfigInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) WithName(name string) *GetProxyConfigInWorkspaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetProxyConfigInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get proxy config in workspace params
func (o *GetProxyConfigInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyConfigInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
