// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteRecipesInWorkspaceReader is a Reader for the DeleteRecipesInWorkspace structure.
type DeleteRecipesInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecipesInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRecipesInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRecipesInWorkspaceOK creates a DeleteRecipesInWorkspaceOK with default headers values
func NewDeleteRecipesInWorkspaceOK() *DeleteRecipesInWorkspaceOK {
	return &DeleteRecipesInWorkspaceOK{}
}

/*DeleteRecipesInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteRecipesInWorkspaceOK struct {
	Payload *model.RecipeV4Responses
}

func (o *DeleteRecipesInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/recipes][%d] deleteRecipesInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteRecipesInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RecipeV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
