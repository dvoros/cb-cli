// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetImageCatalogInWorkspaceReader is a Reader for the GetImageCatalogInWorkspace structure.
type GetImageCatalogInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageCatalogInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImageCatalogInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageCatalogInWorkspaceOK creates a GetImageCatalogInWorkspaceOK with default headers values
func NewGetImageCatalogInWorkspaceOK() *GetImageCatalogInWorkspaceOK {
	return &GetImageCatalogInWorkspaceOK{}
}

/*GetImageCatalogInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetImageCatalogInWorkspaceOK struct {
	Payload *model.ImageCatalogResponse
}

func (o *GetImageCatalogInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/imagecatalogs/{name}][%d] getImageCatalogInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetImageCatalogInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImageCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}