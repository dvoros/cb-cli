package clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPrivateClusterTemplateReader is a Reader for the GetPrivateClusterTemplate structure.
type GetPrivateClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPrivateClusterTemplateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateClusterTemplateOK creates a GetPrivateClusterTemplateOK with default headers values
func NewGetPrivateClusterTemplateOK() *GetPrivateClusterTemplateOK {
	return &GetPrivateClusterTemplateOK{}
}

/*GetPrivateClusterTemplateOK handles this case with default header values.

successful operation
*/
type GetPrivateClusterTemplateOK struct {
	Payload *models_cloudbreak.ClusterTemplateResponse
}

func (o *GetPrivateClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /clustertemplates/user/{name}][%d] getPrivateClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPrivateClusterTemplateOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ClusterTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}