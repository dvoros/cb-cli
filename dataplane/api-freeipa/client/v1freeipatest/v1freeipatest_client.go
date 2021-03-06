// Code generated by go-swagger; DO NOT EDIT.

package v1freeipatest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1freeipatest API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1freeipatest API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UserShowV1 creates free IP a instance
*/
func (a *Client) UserShowV1(params *UserShowV1Params) (*UserShowV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserShowV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userShowV1",
		Method:             "GET",
		PathPattern:        "/v1/freeipa/test/{id}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserShowV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserShowV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
