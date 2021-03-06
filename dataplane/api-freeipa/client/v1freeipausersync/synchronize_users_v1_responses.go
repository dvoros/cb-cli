// Code generated by go-swagger; DO NOT EDIT.

package v1freeipausersync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// SynchronizeUsersV1Reader is a Reader for the SynchronizeUsersV1 structure.
type SynchronizeUsersV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SynchronizeUsersV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSynchronizeUsersV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSynchronizeUsersV1OK creates a SynchronizeUsersV1OK with default headers values
func NewSynchronizeUsersV1OK() *SynchronizeUsersV1OK {
	return &SynchronizeUsersV1OK{}
}

/*SynchronizeUsersV1OK handles this case with default header values.

successful operation
*/
type SynchronizeUsersV1OK struct {
	Payload *model.SynchronizeUsersV1Response
}

func (o *SynchronizeUsersV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/usersync][%d] synchronizeUsersV1OK  %+v", 200, o.Payload)
}

func (o *SynchronizeUsersV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SynchronizeUsersV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
