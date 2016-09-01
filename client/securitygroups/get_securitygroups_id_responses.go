package securitygroups

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

// GetSecuritygroupsIDReader is a Reader for the GetSecuritygroupsID structure.
type GetSecuritygroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSecuritygroupsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecuritygroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecuritygroupsIDOK creates a GetSecuritygroupsIDOK with default headers values
func NewGetSecuritygroupsIDOK() *GetSecuritygroupsIDOK {
	return &GetSecuritygroupsIDOK{}
}

/*GetSecuritygroupsIDOK handles this case with default header values.

successful operation
*/
type GetSecuritygroupsIDOK struct {
	Payload *models.SecurityGroupJSON
}

func (o *GetSecuritygroupsIDOK) Error() string {
	return fmt.Sprintf("[GET /securitygroups/{id}][%d] getSecuritygroupsIdOK  %+v", 200, o.Payload)
}

func (o *GetSecuritygroupsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
