// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StartStackInWorkspaceV4Reader is a Reader for the StartStackInWorkspaceV4 structure.
type StartStackInWorkspaceV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartStackInWorkspaceV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStartStackInWorkspaceV4Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStartStackInWorkspaceV4Default creates a StartStackInWorkspaceV4Default with default headers values
func NewStartStackInWorkspaceV4Default(code int) *StartStackInWorkspaceV4Default {
	return &StartStackInWorkspaceV4Default{
		_statusCode: code,
	}
}

/*StartStackInWorkspaceV4Default handles this case with default header values.

successful operation
*/
type StartStackInWorkspaceV4Default struct {
	_statusCode int
}

// Code gets the status code for the start stack in workspace v4 default response
func (o *StartStackInWorkspaceV4Default) Code() int {
	return o._statusCode
}

func (o *StartStackInWorkspaceV4Default) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/{name}/start][%d] startStackInWorkspaceV4 default ", o._statusCode)
}

func (o *StartStackInWorkspaceV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
