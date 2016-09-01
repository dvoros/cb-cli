package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// GetStacksUserNameReader is a Reader for the GetStacksUserName structure.
type GetStacksUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStacksUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStacksUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStacksUserNameOK creates a GetStacksUserNameOK with default headers values
func NewGetStacksUserNameOK() *GetStacksUserNameOK {
	return &GetStacksUserNameOK{}
}

/*GetStacksUserNameOK handles this case with default header values.

successful operation
*/
type GetStacksUserNameOK struct {
	Payload *models.StackResponse
}

func (o *GetStacksUserNameOK) Error() string {
	return fmt.Sprintf("[GET /stacks/user/{name}][%d] getStacksUserNameOK  %+v", 200, o.Payload)
}

func (o *GetStacksUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
