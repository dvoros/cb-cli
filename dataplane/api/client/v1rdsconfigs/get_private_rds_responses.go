// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPrivateRdsReader is a Reader for the GetPrivateRds structure.
type GetPrivateRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateRdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateRdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateRdsOK creates a GetPrivateRdsOK with default headers values
func NewGetPrivateRdsOK() *GetPrivateRdsOK {
	return &GetPrivateRdsOK{}
}

/*GetPrivateRdsOK handles this case with default header values.

successful operation
*/
type GetPrivateRdsOK struct {
	Payload *model.RDSConfigResponse
}

func (o *GetPrivateRdsOK) Error() string {
	return fmt.Sprintf("[GET /v1/rdsconfigs/user/{name}][%d] getPrivateRdsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateRdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RDSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}