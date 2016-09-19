package cluster

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

// GetConfigsReader is a Reader for the GetConfigs structure.
type GetConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetConfigsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfigsOK creates a GetConfigsOK with default headers values
func NewGetConfigsOK() *GetConfigsOK {
	return &GetConfigsOK{}
}

/*GetConfigsOK handles this case with default header values.

successful operation
*/
type GetConfigsOK struct {
	Payload *models.ConfigsResponse
}

func (o *GetConfigsOK) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/cluster/config][%d] getConfigsOK  %+v", 200, o.Payload)
}

func (o *GetConfigsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
