// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClusterReader is a Reader for the DeleteCluster structure.
type DeleteClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteClusterDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteClusterDefault creates a DeleteClusterDefault with default headers values
func NewDeleteClusterDefault(code int) *DeleteClusterDefault {
	return &DeleteClusterDefault{
		_statusCode: code,
	}
}

/*DeleteClusterDefault handles this case with default header values.

successful operation
*/
type DeleteClusterDefault struct {
	_statusCode int
}

// Code gets the status code for the delete cluster default response
func (o *DeleteClusterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/stacks/{id}/cluster][%d] deleteCluster default ", o._statusCode)
}

func (o *DeleteClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
