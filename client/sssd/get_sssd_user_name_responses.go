package sssd

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

// GetSssdUserNameReader is a Reader for the GetSssdUserName structure.
type GetSssdUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSssdUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSssdUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSssdUserNameOK creates a GetSssdUserNameOK with default headers values
func NewGetSssdUserNameOK() *GetSssdUserNameOK {
	return &GetSssdUserNameOK{}
}

/*GetSssdUserNameOK handles this case with default header values.

successful operation
*/
type GetSssdUserNameOK struct {
	Payload *models.SssdConfigResponse
}

func (o *GetSssdUserNameOK) Error() string {
	return fmt.Sprintf("[GET /sssd/user/{name}][%d] getSssdUserNameOK  %+v", 200, o.Payload)
}

func (o *GetSssdUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SssdConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
