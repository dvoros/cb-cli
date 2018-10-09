// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
)

// PostStackForBlueprintV3Reader is a Reader for the PostStackForBlueprintV3 structure.
type PostStackForBlueprintV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStackForBlueprintV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostStackForBlueprintV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStackForBlueprintV3OK creates a PostStackForBlueprintV3OK with default headers values
func NewPostStackForBlueprintV3OK() *PostStackForBlueprintV3OK {
	return &PostStackForBlueprintV3OK{}
}

/*PostStackForBlueprintV3OK handles this case with default header values.

successful operation
*/
type PostStackForBlueprintV3OK struct {
	Payload *model.GeneratedBlueprintResponse
}

func (o *PostStackForBlueprintV3OK) Error() string {
	return fmt.Sprintf("[POST /v3/{workspaceId}/stacks/{name}/blueprint][%d] postStackForBlueprintV3OK  %+v", 200, o.Payload)
}

func (o *PostStackForBlueprintV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GeneratedBlueprintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}