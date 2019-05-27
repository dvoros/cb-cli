// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetRegionsByCredentialAndWorkspaceReader is a Reader for the GetRegionsByCredentialAndWorkspace structure.
type GetRegionsByCredentialAndWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsByCredentialAndWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsByCredentialAndWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegionsByCredentialAndWorkspaceOK creates a GetRegionsByCredentialAndWorkspaceOK with default headers values
func NewGetRegionsByCredentialAndWorkspaceOK() *GetRegionsByCredentialAndWorkspaceOK {
	return &GetRegionsByCredentialAndWorkspaceOK{}
}

/*GetRegionsByCredentialAndWorkspaceOK handles this case with default header values.

successful operation
*/
type GetRegionsByCredentialAndWorkspaceOK struct {
	Payload *model.CompactRegionV1Response
}

func (o *GetRegionsByCredentialAndWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v1/platform_resources/regions][%d] getRegionsByCredentialAndWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetRegionsByCredentialAndWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CompactRegionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
