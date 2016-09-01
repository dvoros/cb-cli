package networks

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

// PostNetworksAccountReader is a Reader for the PostNetworksAccount structure.
type PostNetworksAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostNetworksAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNetworksAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNetworksAccountOK creates a PostNetworksAccountOK with default headers values
func NewPostNetworksAccountOK() *PostNetworksAccountOK {
	return &PostNetworksAccountOK{}
}

/*PostNetworksAccountOK handles this case with default header values.

successful operation
*/
type PostNetworksAccountOK struct {
	Payload *models.ID
}

func (o *PostNetworksAccountOK) Error() string {
	return fmt.Sprintf("[POST /networks/account][%d] postNetworksAccountOK  %+v", 200, o.Payload)
}

func (o *PostNetworksAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
