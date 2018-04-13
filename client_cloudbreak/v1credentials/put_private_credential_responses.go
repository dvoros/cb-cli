// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PutPrivateCredentialReader is a Reader for the PutPrivateCredential structure.
type PutPrivateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPrivateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutPrivateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPrivateCredentialOK creates a PutPrivateCredentialOK with default headers values
func NewPutPrivateCredentialOK() *PutPrivateCredentialOK {
	return &PutPrivateCredentialOK{}
}

/*PutPrivateCredentialOK handles this case with default header values.

successful operation
*/
type PutPrivateCredentialOK struct {
	Payload *models_cloudbreak.CredentialResponse
}

func (o *PutPrivateCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /v1/credentials/user][%d] putPrivateCredentialOK  %+v", 200, o.Payload)
}

func (o *PutPrivateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.CredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}