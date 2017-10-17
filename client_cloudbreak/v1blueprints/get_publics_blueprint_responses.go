// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicsBlueprintReader is a Reader for the GetPublicsBlueprint structure.
type GetPublicsBlueprintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsBlueprintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsBlueprintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsBlueprintOK creates a GetPublicsBlueprintOK with default headers values
func NewGetPublicsBlueprintOK() *GetPublicsBlueprintOK {
	return &GetPublicsBlueprintOK{}
}

/*GetPublicsBlueprintOK handles this case with default header values.

successful operation
*/
type GetPublicsBlueprintOK struct {
	Payload []*models_cloudbreak.BlueprintResponse
}

func (o *GetPublicsBlueprintOK) Error() string {
	return fmt.Sprintf("[GET /v1/blueprints/account][%d] getPublicsBlueprintOK  %+v", 200, o.Payload)
}

func (o *GetPublicsBlueprintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
