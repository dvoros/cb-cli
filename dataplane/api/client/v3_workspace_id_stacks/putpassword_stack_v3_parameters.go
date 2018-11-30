// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutpasswordStackV3Params creates a new PutpasswordStackV3Params object
// with the default values initialized.
func NewPutpasswordStackV3Params() *PutpasswordStackV3Params {
	var ()
	return &PutpasswordStackV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutpasswordStackV3ParamsWithTimeout creates a new PutpasswordStackV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutpasswordStackV3ParamsWithTimeout(timeout time.Duration) *PutpasswordStackV3Params {
	var ()
	return &PutpasswordStackV3Params{

		timeout: timeout,
	}
}

// NewPutpasswordStackV3ParamsWithContext creates a new PutpasswordStackV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutpasswordStackV3ParamsWithContext(ctx context.Context) *PutpasswordStackV3Params {
	var ()
	return &PutpasswordStackV3Params{

		Context: ctx,
	}
}

// NewPutpasswordStackV3ParamsWithHTTPClient creates a new PutpasswordStackV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutpasswordStackV3ParamsWithHTTPClient(client *http.Client) *PutpasswordStackV3Params {
	var ()
	return &PutpasswordStackV3Params{
		HTTPClient: client,
	}
}

/*PutpasswordStackV3Params contains all the parameters to send to the API endpoint
for the putpassword stack v3 operation typically these are written to a http.Request
*/
type PutpasswordStackV3Params struct {

	/*Body*/
	Body *model.UserNamePassword
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithTimeout(timeout time.Duration) *PutpasswordStackV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithContext(ctx context.Context) *PutpasswordStackV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithHTTPClient(client *http.Client) *PutpasswordStackV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithBody(body *model.UserNamePassword) *PutpasswordStackV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetBody(body *model.UserNamePassword) {
	o.Body = body
}

// WithName adds the name to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithName(name string) *PutpasswordStackV3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) WithWorkspaceID(workspaceID int64) *PutpasswordStackV3Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the putpassword stack v3 params
func (o *PutpasswordStackV3Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutpasswordStackV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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