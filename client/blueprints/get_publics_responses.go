package blueprints

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

// GetPublicsReader is a Reader for the GetPublics structure.
type GetPublicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPublicsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsOK creates a GetPublicsOK with default headers values
func NewGetPublicsOK() *GetPublicsOK {
	return &GetPublicsOK{}
}

/*GetPublicsOK handles this case with default header values.

successful operation
*/
type GetPublicsOK struct {
	Payload []*models.BlueprintResponse
}

func (o *GetPublicsOK) Error() string {
	return fmt.Sprintf("[GET /blueprints/account][%d] getPublicsOK  %+v", 200, o.Payload)
}

func (o *GetPublicsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
