// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// GetAllTenantsReader is a Reader for the GetAllTenants structure.
type GetAllTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllTenantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAllTenantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllTenantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllTenantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllTenantsOK creates a GetAllTenantsOK with default headers values
func NewGetAllTenantsOK() *GetAllTenantsOK {
	return &GetAllTenantsOK{}
}

/*GetAllTenantsOK handles this case with default header values.

successful operation
*/
type GetAllTenantsOK struct {
	Payload []*model.Tenant
}

func (o *GetAllTenantsOK) Error() string {
	return fmt.Sprintf("[GET /caas/api/tenants][%d] getAllTenantsOK  %+v", 200, o.Payload)
}

func (o *GetAllTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTenantsBadRequest creates a GetAllTenantsBadRequest with default headers values
func NewGetAllTenantsBadRequest() *GetAllTenantsBadRequest {
	return &GetAllTenantsBadRequest{}
}

/*GetAllTenantsBadRequest handles this case with default header values.

Bad Request
*/
type GetAllTenantsBadRequest struct {
}

func (o *GetAllTenantsBadRequest) Error() string {
	return fmt.Sprintf("[GET /caas/api/tenants][%d] getAllTenantsBadRequest ", 400)
}

func (o *GetAllTenantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsUnauthorized creates a GetAllTenantsUnauthorized with default headers values
func NewGetAllTenantsUnauthorized() *GetAllTenantsUnauthorized {
	return &GetAllTenantsUnauthorized{}
}

/*GetAllTenantsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllTenantsUnauthorized struct {
}

func (o *GetAllTenantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /caas/api/tenants][%d] getAllTenantsUnauthorized ", 401)
}

func (o *GetAllTenantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsForbidden creates a GetAllTenantsForbidden with default headers values
func NewGetAllTenantsForbidden() *GetAllTenantsForbidden {
	return &GetAllTenantsForbidden{}
}

/*GetAllTenantsForbidden handles this case with default header values.

Forbidden
*/
type GetAllTenantsForbidden struct {
}

func (o *GetAllTenantsForbidden) Error() string {
	return fmt.Sprintf("[GET /caas/api/tenants][%d] getAllTenantsForbidden ", 403)
}

func (o *GetAllTenantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsInternalServerError creates a GetAllTenantsInternalServerError with default headers values
func NewGetAllTenantsInternalServerError() *GetAllTenantsInternalServerError {
	return &GetAllTenantsInternalServerError{}
}

/*GetAllTenantsInternalServerError handles this case with default header values.

Internal server error
*/
type GetAllTenantsInternalServerError struct {
}

func (o *GetAllTenantsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /caas/api/tenants][%d] getAllTenantsInternalServerError ", 500)
}

func (o *GetAllTenantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
