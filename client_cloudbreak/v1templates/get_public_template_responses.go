// Code generated by go-swagger; DO NOT EDIT.

package v1templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicTemplateReader is a Reader for the GetPublicTemplate structure.
type GetPublicTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicTemplateOK creates a GetPublicTemplateOK with default headers values
func NewGetPublicTemplateOK() *GetPublicTemplateOK {
	return &GetPublicTemplateOK{}
}

/*GetPublicTemplateOK handles this case with default header values.

successful operation
*/
type GetPublicTemplateOK struct {
	Payload *models_cloudbreak.TemplateResponse
}

func (o *GetPublicTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v1/templates/account/{name}][%d] getPublicTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPublicTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.TemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
